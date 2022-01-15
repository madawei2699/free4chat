<template>
  <div class="normal-page-layout">
    <slot name="app-bar">
      <f-app-bar
        dark
        @go-home="goHome"
      >
        <v-btn
          small
          icon
          class="ml-0 mr-0"
          :color="'#ffffff'"
          @click="openProfile"
        >
          <v-icon>{{ $icons.mdiAccountCircleOutline }}</v-icon>
        </v-btn>
      </f-app-bar>
    </slot>
    <slot />
    <v-dialog
      v-model="showProfileDialog"
      max-width="290"
      transition="dialog-bottom-transition"
      fullscreen
      hide-overlay
    >
      <v-card color="#000000">
        <v-toolbar dark elevation="0" color="#000000">
          <v-toolbar-title>Profile</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn icon @click="close">
              <v-icon>
                {{ $icons.mdiClose }}
              </v-icon>
            </v-btn>
          </v-toolbar-items>
        </v-toolbar>

        <v-card-text class="mb-0 pl-4 pr-5">
          <v-list-item class="px-0">
            <v-list-item-content>
              <v-list-item-subtitle class="mb-1 overline">My User ID</v-list-item-subtitle>
              <v-list-item-title class="body-2">
                <span :style="uuidColor">{{ profile.uid }}</span>
              </v-list-item-title>
            </v-list-item-content>

            <v-list-item-action>
              <v-btn icon @click="resetUid">
                <v-icon small>
                  {{ $icons.mdiRefresh }}
                </v-icon>
              </v-btn>
            </v-list-item-action>
          </v-list-item>
        </v-card-text>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
import { State, Mutation } from 'vuex-class'
import { uuidToColor, uuidv4 } from '@/utils/uuid'

@Component
class NormalPageLayout extends Vue {
  @Prop() hasBack!: boolean
  @State(state => state.app.profile) profile
  @Mutation('app/SET_PROFILE') setProfile

  showProfileDialog:boolean = false

  get uuidColor () {
    const gb = uuidToColor(this.profile.uid)
    const g = gb[0]
    const b = gb[1]
    const r = 180
    let color = ''
    if (g % 3 === 0) {
      color = `rgb(${r},${g},${b})`
    } else if (g % 3 === 1) {
      color = `rgb(${g},${r},${b})`
    } else {
      color = `rgb(${g},${b},${r})`
    }
    return { color }
  }

  back () {
    this.$router.back()
  }

  resetUid () {
    const uid = uuidv4()
    this.setProfile({ uid })
  }

  close () {
    this.showProfileDialog = false
  }

  openProfile () {
    this.showProfileDialog = true
  }

  goHome () {
    console.log('normal layout go home')
    this.$emit('home')
  }
}
export default NormalPageLayout
</script>
<style lang="scss" scoped>
.page-nav {
  padding: 0 8px;

  .page-title {
    font-size: 22px;
    color: #000000;
  }
}

.normal-page-layout {
  height: 100%;
  display: flex;
  flex-direction: column;
  background-size: cover;
  z-index: 2;
}

</style>
