<template>
  <v-app-bar
    v-if="show"
    flat
    height="48"
    :dark="dark"
    :color="color"
    :style="style"
    class="appbar"
  >
    <v-btn
      v-if="hasBack"
      small
      icon
      class="ml-0"
      :color="dark ? '#ffffff' : '#111111'"
      @click="handleBack"
    >
      <v-icon>{{ $icons.mdiArrowLeft }}</v-icon>
    </v-btn>
    <v-btn
      v-if="hasHome"
      small
      icon
      class="ml-0 mr-2"
      :color="dark ? '#ffffff' : '#111111'"
      @click="handleHome"
    >
      <v-icon>{{ $icons.mdiHome }}</v-icon>
    </v-btn>
    <v-toolbar-title class="pl-2 font-weight-bold">
      {{ appbar.title }}
    </v-toolbar-title>
    <v-spacer />
    <slot />
  </v-app-bar>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { State } from 'vuex-class'

@Component
class FAppBar extends Vue {
  @State(state => state.app.appbar) appbar

  get hasBack () {
    return this.appbar.back
  }

  get hasHome () {
    return this.appbar.home
  }

  get show () {
    return this.appbar.show
  }

  get color () {
    return this.appbar.color
  }

  get dark () {
    return this.appbar.dark
  }

  get style () {
    return this.appbar.style
  }

  handleBack () {
    this.$router.back()
  }

  handleHome () {
    this.$emit('go-home')
    this.$router.push('/')
  }
}
export default FAppBar
</script>

<style lang="scss" scoped>
.appbar {
  flex: 0;
  background: linear-gradient(to bottom, #040c11 0%, #040c1100 100%) !important;
}
</style>
