<template>
  <normal-page-layout
    @home="goHome"
  >
    <v-container class="room-page">
      <f-loading :loading="loading" :fullscreen="true" />
      <template>
        <div v-if="isEmpty" class="hint-box empty-hint px-4 py-2 mx-2 mb-4">
          <h2 class="body-1 font-weight-bold">{{ $t('room.empty_block_title') }}</h2>
          <div class="body-2">
            {{ $t('room.empty_block_text') }}
          </div>
          <v-btn
            class="my-2"
            small
            outlined
            v-clipboard:copy="destination"
            v-clipboard:success="() => copyUtil.onCopy(this)"
            v-clipboard:error="() => copyUtil.onError(this)"
          >
            {{ $t('common.copy_url') }}
          </v-btn>
        </div>
        <div v-else class="hint-box caption px-4 py-2 mx-2 mb-4">
          <v-icon small>{{ $icons.mdiHeadphones }}</v-icon>
          {{ $t('room.earhub_block_text') }}
        </div>
        <div v-if="noMicPermission" class="hint-box error-hint px-4 py-2 mx-2 mb-4">
          <h2 class="body-1 font-weight-bold">{{ $t('room.error_block_title') }}</h2>
          <div class="body-2">
            {{ $t('room.error_block_text') }}
          </div>
        </div>
        <div v-if="isWeChat" class="hint-box error-hint px-4 py-2 mx-2 mb-4">
          <h2 class="body-1 font-weight-bold">{{ $t('room.wechat_error_block_title') }}</h2>
          <div class="body-2">
            {{ $t('room.wechat_error_block_text') }}
          </div>
        </div>
        <div class="cards">
          <div
            v-for="user in participants"
            :key="user.uid"
            :style="{ 'width': cardWidth }"
            class="user-card-wrapper ma-2"
          >
            <user-card
              :user="user"
              @mute="muteUser"
              @unmute="unmuteUser"
              @edit-name="editName"
            />
          </div>
        </div>
      </template>
    </v-container>
    <v-dialog
      v-model="showNameDialog"
      max-width="290"
    >
      <v-card>
        <v-card-title class="title-2">{{ $t('room.name_dialog_title') }}</v-card-title>
        <v-card-text class="mb-0">
          <div class="mb-4">
            <v-text-field
              v-model="nickname"
              :label="$t('room.name_dialog_placeholder')"
              :hide-details="true"
            >
              <template v-slot:append>
                <v-btn icon @click="genRandomChannelName">
                  <v-icon>
                    {{ $icons.mdiDice3 }}
                  </v-icon>
                </v-btn>
              </template>
            </v-text-field>
          </div>
        </v-card-text>
        <v-card-actions class="px-5 pb-5">
          <v-btn
            color="primary"
            block
            rounded
            :disabled="!validated"
            @click="join"
          >
            {{ $t('room.name_dialog_btn') }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-btn
      color="rgba(255, 255, 255, 0.2)"
      dark
      small
      absolute
      bottom
      right
      fab
      class="mute-all-btn"
      @click="muteAll"
    >
      <v-icon :color="this.isAllMuted ? 'rgb(242, 72, 34)' : '#FFFFFF'">
        {{ this.isAllMuted ? $icons.mdiMicrophoneOff : $icons.mdiMicrophone }}
      </v-icon>
    </v-btn>
  </normal-page-layout>
</template>

<script lang="ts">
import { Component, Mixins } from 'vue-property-decorator'
import { Mutation, State } from 'vuex-class'
import { uniqueNamesGenerator, Config, adjectives, starWars } from 'unique-names-generator'
import PageView from '@/mixins/page'
import { launch, stop } from '@/services/rpc'
import { uuidv4 } from '@/utils/uuid'
import copyUtil from '@/utils/copy'
import UserCard from '@/components/UserCard.vue'

const randomNameConfig: Config = {
  dictionaries: [adjectives, starWars],
  separator: ' ',
  length: 2
}

@Component({
  head () {
    return {
      title: this.title,
      meta: [
        {
          hid: 'theme-color',
          name: 'theme-color',
          content: '#040C11'
        }
      ]
    }
  },
  components: {
    UserCard
  }
})
class RoomPage extends Mixins(PageView) {
  @State(state => state.app.profile) profile
  @State(state => state.app.chat) chat
  @Mutation('app/SET_NICKNAME') setNickname
  @Mutation('app/SET_PROFILE') setProfile
  @Mutation('app/SET_APPBAR') setAppbar
  @Mutation('app/ADD_MUTE') addMute
  @Mutation('app/REMOVE_MUTE') removeMute
  @Mutation('app/ADD_MUTES') addMutes
  @Mutation('app/REMOVE_MUTES') removeMutes

  loading = false

  showNameDialog = false

  noMicPermission:boolean = false

  nickname = ''

  roomName = ''

  uid = ''

  participantMap = {}

  participantTrackIdMap = {}

  participants:any = []

  streams = {}

  pc:any = null

  copyUtil = copyUtil

  get isAllMuted () {
    for (let ix = 0; ix < this.participants.length; ix++) {
      if (!this.participants[ix].isMuted) {
        return false
      }
    }
    return true
  }

  get isWeChat () {
    // MicroMessenger
    return navigator.userAgent.toLowerCase().includes('micromessenger')
  }

  get destination () {
    const url = (window as any).location.href
    return url
  }

  get validated () {
    return this.nickname.trim().length !== 0
  }

  get cardWidth () {
    const winWidth = window.innerWidth
    if (winWidth < 460) {
      const cw = Math.round((winWidth - 12 * 2 - 10 * 4) / 2)
      return `${cw}px`
    }
    return '200px'
  }

  get isEmpty () {
    return this.participants.length === 1
  }

  get title () {
    return `#${this.roomName}`
  }

  genRandomChannelName () {
    const name: string = uniqueNamesGenerator(randomNameConfig)
    this.nickname = name.slice(0, 1).toUpperCase() + name.slice(1)
  }

  mounted () {
    setTimeout(() => {
      this.reload()
      if (this.nickname) {
        launch(this.roomName, this.nickname, this.uid, this.onConnect, this.onDisconnect, this.onResume, this.onError)
      }
    }, 100)
  }

  reload () {
    this.loading = true
    const rname = this.$route.params.name

    if (!this.chat.rooms.hasOwnProperty(rname)) {
      this.showNameDialog = true
    } else {
      this.nickname = this.chat.rooms[rname].nickname
    }
    this.uid = this.profile.uid || uuidv4()
    this.setProfile({ uid: this.uid })

    this.roomName = rname
    this.setAppbar({
      color: 'rgba(0, 0, 0, 0.0)',
      home: true,
      back: false,
      title: `#${this.roomName}`
    })
    this.loading = false
  }

  join () {
    this.setNickname({ room: this.roomName, nickname: this.nickname })
    this.showNameDialog = false
    if (this.participants.length !== 0) {
      window.location.reload()
      return
    }
    launch(this.roomName, this.nickname, this.uid, this.onConnect, this.onDisconnect, this.onResume, this.onError)
  }

  onConnect (pc:any, stream:any, analyser:any, trackId:string, targetUid:string, targetNickname:string) {
    this.pc = pc
    console.log(stream, targetUid, targetNickname)
    this.addParticipant(stream, analyser, trackId, targetUid, targetNickname)
    this.streams[targetUid] = stream

    const defaultMuteValue = this.chat.mutes.hasOwnProperty(targetUid)
    if (defaultMuteValue) {
      this.muteOrUnmute(targetUid, true)
    }
  }

  onDisconnect (trackId:string) {
    const user:any = this.removeParticipant(trackId)
    if (user !== null) {
      delete this.streams[user.uid]
    }
  }

  onError (err:any) {
    // constraint: ""
    // message: "The request is not allowed by the user agent or the platform in the current context."
    // name: "NotAllowedError"
    // stack: ""
    if (err && err.name === 'NotAllowedError') {
      console.log('no permission')
      this.noMicPermission = true
    } else {
      console.log(err)
    }
  }

  onResume (err:any) {
    return new Promise((resolve) => {
      console.log('resume from connection', err)
      this.clearup()
      resolve()
    })
  }

  clearup () {
    console.log('clear up')
    this.participants.splice(0, this.participants.length)
    this.participantTrackIdMap = {}
    this.participantMap = {}
    this.streams = {}
  }

  goHome () {
    stop()
    this.clearup()
  }

  test () {
    console.log(this.participantTrackIdMap)
    console.log(this.participantMap)
    console.log(this.participants)
  }

  findParticipant (uid) {
    for (let ix = 0; ix < this.participants.length; ix++) {
      if (this.participants[ix].uid === uid) {
        return { user: this.participants[ix], index: ix }
      }
    }
    return null
  }

  addParticipant (stream:any, analyser:any, trackId:string, targetUid:string, targetNickname:string) {
    const incoming:any = {
      stream,
      analyser,
      trackId,
      isMuted: false,
      uid: targetUid,
      nickname: targetNickname
    }
    if (!this.participantMap.hasOwnProperty(targetUid)) {
      this.participants.push(incoming)
    }
    console.log('addParticipant', targetUid, targetNickname, trackId)
    this.participantMap[targetUid] = incoming
    this.participantTrackIdMap[trackId] = incoming
  }

  removeParticipant (trackId:string) {
    if (!this.participantTrackIdMap.hasOwnProperty(trackId)) {
      return null
    }
    const user:any = this.participantTrackIdMap[trackId]
    if (user) {
      const ret:any = this.findParticipant(user.uid)
      if (ret && ret.user) {
        console.log('removeParticipant', user, trackId)
        this.participants.splice(ret.index, 1)
        delete this.participantMap[user.uid]
      }
    }
    delete this.participantTrackIdMap[trackId]
    return user
  }

  editName (user) {
    if (user && this.uid === user.uid) {
      this.nickname = user.nickname
      this.showNameDialog = true
    }
  }

  muteAll () {
    const userIds = this.participants.map((x) => {
      return x.uid
    })
    const value = this.isAllMuted
    for (let ix = 0; ix < userIds.length; ix++) {
      this.muteOrUnmute(userIds[ix], !value)
    }
    if (value) {
      this.removeMutes(userIds)
    } else {
      this.addMutes(userIds)
    }
  }

  muteUser (user) {
    this.muteOrUnmute(user.uid, true)
    this.addMute(user.uid)
  }

  unmuteUser (user) {
    this.muteOrUnmute(user.uid, false)
    this.removeMute(user.uid)
  }

  muteOrUnmute (uid, value) {
    if (!this.participantMap.hasOwnProperty(uid)) {
      return
    }
    console.log(`uid=${uid}, mute=${value}`)
    const ret:any = this.findParticipant(uid)
    if (ret && ret.user) {
      ret.user.isMuted = value
      this.participants.splice(ret.index, 1, ret.user)
      this.setStream(uid, !value)
    }
  }

  setStream (uid, val) {
    const stream = this.streams[uid]
    if (stream && stream.getTracks().length > 0) {
      stream.getTracks()[0].enabled = val
    }
  }
}
export default RoomPage
</script>

<style lang="scss" scoped>
.mute-all-btn {
  bottom: 0;
  position: absolute;
  margin: 0 16px 68px 0;
  box-shadow: 0 0 20px rgba(255, 255, 255, 0.08);
}
.cards {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  .user-card-wrapper {
  }
}
.hint-box {
  color: rgba(255, 255, 255, 0.7);
  border-radius: 20px;
  background: linear-gradient(-15deg, rgba(255, 255, 255, 0.12) 0%, rgba(255, 255, 255, 0.1) 100%);
  &.empty-hint {
    background: linear-gradient(-15deg, rgb(153, 0, 255) 0%, #0057fa 100%) !important;
  }
  &.error-hint {
    background: linear-gradient(-15deg, rgb(236, 87, 0) 0%, #be1d00 100%) !important;
  }
}
</style>
