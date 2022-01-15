<template>
  <v-app class="app">
    <nuxt />
    <toast />
    <ul v-if="useBgAnimation" class="bg-bubbles" :class="useBgAnimation ? '' : 'pause'">
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
      <li></li>
    </ul>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator'
import { State } from 'vuex-class'

@Component({
  middleware: 'i18n'
})
class DefaultLayout extends Vue {
  @State(state => state.app.profile) profile
  @State(state => state.app.appbar) appbar

  useBgAnimation = false

  mounted () {
    (window as any).onNuxtReady(() => {
      if (this.appbar.animation) {
        this.useBgAnimation = true
      }
    })
  }

  @Watch('appbar.animation')
  onAnimationChange () {
    this.useBgAnimation = this.appbar.animation
  }
}
export default DefaultLayout
</script>

<style lang="scss" scoped>
.v-application--is-ltr .v-toolbar__content > .v-btn.v-btn--icon:first-child + .v-toolbar__title, .v-application--is-ltr .v-toolbar__extension > .v-btn.v-btn--icon:first-child + .v-toolbar__title {
  padding-left: 8px;
}

.app {
  overflow: hidden;
  background: #2e2e2e;
  background: linear-gradient(to bottom right, #040c11 0%, #002950 100%) !important;
}

.bg-bubbles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1;
  &.pause {
    li {
      animation-play-state: paused !important;
    }
  }
  li {
    position: absolute;
    list-style: none;
    display: block;
    width: 40px;
    height: 40px;
    background-color: rgba(255, 255, 255, 0.05);
    bottom: -160px;

    animation-duration: 133s;

    -webkit-animation: square 25s infinite;
    animation:         square 25s infinite;

    -webkit-transition-timing-function: linear;
    transition-timing-function: linear;

    &:nth-child(1){
      left: 10%;
    }
    &:nth-child(2){
      left: 20%;
      width: 80px;
      height: 80px;
      animation-delay: -20s;
      animation-duration: 97s;
    }
    &:nth-child(3){
      left: 25%;
      animation-delay: -10s;
    }
    &:nth-child(4){
      left: 40%;
      width: 60px;
      height: 60px;
      animation-duration: 102s;
      animation-delay: -20s;
      background-color: rgba(255, 255, 255, 0.15);
    }
    &:nth-child(5){
      animation-delay: -30s;
      left: 70%;
    }
    &:nth-child(6){
      left: 80%;
      width: 120px;
      height: 120px;
      animation-delay: -13s;
      background-color: rgba(255, 255, 255, 0.08);
    }
    &:nth-child(7){
      left: 32%;
      width: 160px;
      height: 160px;
      animation-delay: -31s;
    }
    &:nth-child(8){
      left: 55%;
      width: 20px;
      height: 20px;
      animation-delay: -65s;
      animation-duration: 140s;
    }
    &:nth-child(9){
      left: 25%;
      width: 10px;
      height: 10px;
      animation-delay: -2s;
      animation-duration: 180s;
      background-color: rgba(255, 255, 255, 0.12);
    }
    &:nth-child(10){
      left: 90%;
      width: 160px;
      height: 160px;
      animation-delay: -11s;
    }
  }
}
@-webkit-keyframes square {
  0%   { transform: translateY(0); }
  100% { transform: translateY(-2000px) rotate(600deg); }
}
@keyframes square {
  0%   { transform: translateY(0); }
  100% { transform: translateY(-2000px) rotate(600deg); }
}
</style>
