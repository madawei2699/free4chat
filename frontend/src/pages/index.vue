<template>
  <normal-page-layout>
    <v-container class="room-page">
      <f-loading :loading="loading" :fullscreen="true" />
      <template>
        <div class="intro pa-4">
          <p class="display-1 font-weight-bold">{{ $t('index.headline_1') }}</p>
          <p class="title font-weight-normal">{{ $t('index.headline_2') }}</p>
          <div class="d-flex py-5 start-field-wrapper">
            <v-text-field
              prefix="#"
              label="Room Name"
              class="room-field"
              solo
              v-model="channelName"
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
            <v-btn
              large
              class="start-btn"
              color="primary"
              rounded
              :disabled="!validated"
              @click="createOrJoin"
            >
              {{ $t('index.create_room_btn') }}
            </v-btn>
          </div>
          <div class="hint">
            <p class="caption" v-html="$t('index.create_room_dialog_hint_1')"></p>
            <p class="caption" v-html="$t('index.create_room_dialog_hint_2')"></p>
          </div>
        </div>
        <template v-if="!noRecentRooms">
          <v-subheader class="body-2 ml-2">{{ $t('index.recent_room_label') }}</v-subheader>
          <div class="rooms">
            <div
              v-for="room in recentRooms"
              :key="`${room.room}-${room.nickname}`"
              :style="{ 'width': cardWidth }"
              class="rooms-card-wrapper ma-2"
            >
              <room-card
                :room="room"
                @join="joinRoom"
                @delete="deleteRoom"
              >
              </room-card>
            </div>
          </div>
        </template>
      </template>
      <div class="faq px-4">
        <div class="faq-title display-1 mt-4 py-4">{{ $t('index.faq_title') }}</div>
        <div
          v-for="item in faqItems"
          :key="item.q"
          class="faq-item"
        >
          <div class="question font-weight-bold" v-html="item.q"></div>
          <div class="anwser" v-html="item.a"></div>
        </div>
      </div>
      <div class="footer">
        <div class="twitter">
          © Mornin 2020<br>
          Follow <a href="https://twitter.com/MorninFM">@MorninFM</a> to contact us.
        </div>
      </div>
    </v-container>
  </normal-page-layout>
</template>

<script lang="ts">
import { Component, Mixins } from 'vue-property-decorator'
import { Mutation, State } from 'vuex-class'
import { uniqueNamesGenerator, Config, adjectives, countries } from 'unique-names-generator'
import PageView from '@/mixins/page'
import RoomCard from '@/components/RoomCard.vue'

const randomNameConfig: Config = {
  dictionaries: [adjectives, countries],
  separator: '-',
  length: 2
}

@Component({
  head () {
    return {
      title: this.title
    }
  },
  components: {
    RoomCard
  }
})
class IndexPage extends Mixins(PageView) {
  @State(state => state.app.profile) profile
  @State(state => state.app.chat) chat
  @Mutation('app/SET_APPBAR') setAppbar
  @Mutation('app/REMOVE_ROOM') removeRoom

  showJoinDialog = false

  loading = false

  channelName:string = ''

  faqItems:any = []

  get title () {
    return 'Mornin'
  }

  get cardWidth () {
    const winWidth = window.innerWidth
    if (winWidth < 600) {
      return `${winWidth - 12 - 8 * 2}px`
    }
    return '240px'
  }

  get recentRooms () {
    return this.chat.rooms
  }

  get noRecentRooms () {
    return Object.values(this.recentRooms).length === 0
  }

  get validated () {
    return this.channelName.trim().length !== 0 && /^[a-zA-Z0-9\-_]+$/.test(this.channelName)
  }

  mounted () {
    for (let ix = 0; ix < 9; ix++) {
      this.faqItems.push({
        q: this.$i18n.t(`faq.q${ix + 1}`),
        a: this.$i18n.t(`faq.a${ix + 1}`)
      })
    }
    setTimeout(() => {
      this.reload()
    }, 100)
  }

  reload () {
    this.setAppbar({
      color: 'rgba(0,0,0,0.0)',
      title: 'Mornin',
      animation: true,
      back: false
    })
  }

  genRandomChannelName () {
    const name:string = uniqueNamesGenerator(randomNameConfig)
    this.channelName = name.toLowerCase().replace(/[^a-zA-Z0-9\-_]/g, '')
  }

  deleteRoom (room) {
    this.removeRoom({ room: room.room })
  }

  joinRoom (room) {
    this.$router.push(`/${room.room}`)
  }

  createOrJoin () {
    this.$router.push(`/${this.channelName}`)
  }
}
export default IndexPage
</script>

<style lang="scss" scoped>
.intro {
  margin-top: 40px;
  .start-field-wrapper {
    align-items: center;
    margin-top: 40px;
    max-width: 600px;
    ::v-deep .v-input__append-inner {
      margin-top: 2px;
    }
    .room-field {
      margin-right: 16px;
    }
    .start-btn {
      min-width: 180px;
    }
  }
}

.hint {
  color: rgba(255, 255, 255, 0.7);
  p {
    line-height: 1.3;
    margin-bottom: 1em;
  }
}
.domain {
  color: #2196f3;
}
.rooms {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  .room-card-wrapper {
  }
}

@media only screen and (max-width: 736px) {
  .start-field-wrapper {
    width: 100%;
    align-items: flex-start !important;
    justify-content: flex-start;
    flex-direction: column;
    ::v-deep .v-input__append-inner {
      margin-top: 2px;
    }
    .room-field {
      width: 100%;
    }
    .start-btn {
      margin-top: 14px;
    }
  }
}

.faq {
  margin-top: 40px;
  padding-right: 20px !important;
  .faq-item {
    margin: 0px 0 20px 0;
    opacity: 0.8;
    .question {
      margin: 4px 0;
    }
    .anwser {
      margin: 4px 0;
    }
  }
}

.footer {
  color: rgba(255, 255, 255, 0.7);
  margin-top: 60px;
  padding: 16px;
}
</style>
