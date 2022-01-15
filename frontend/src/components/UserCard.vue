<template>
  <div
    v-if="user"
    class="user-card"
  >
    <div class="top d-flex py-2 pl-4 pr-2">
      <div class="nickname font-weight-bold" @click="editName">
        {{ user.nickname }}
      </div>
      <v-btn
        icon
        small
        @click="handleMutable"
      >
        <v-icon :color="this.user.isMuted ? 'rgb(242, 72, 34)' : '#FFFFFF'">
          {{ this.user.isMuted ? $icons.mdiMicrophoneOff : $icons.mdiMicrophone }}
        </v-icon>
      </v-btn>
    </div>
    <div class="canvas-wrapper">
      <canvas class="canvas" ref="canvas" />
    </div>
    <audio ref="audio" autoplay />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
import { uuidToColor } from '@/utils/uuid'

@Component
class FLoading extends Vue {
  @Prop() user
  @Prop() muted

  mounted () {
    setTimeout(() => {
      if (this.user.trackId !== 'me') {
        const audio:any = (this.$refs.audio as any)
        audio.srcObject = this.user.stream
      }
      this.visualize(this.user.uid, this.user.analyser)
    }, 200)
  }

  editName () {
    this.$emit('edit-name', this.user)
  }

  handleMutable () {
    if (this.user.isMuted) {
      this.$emit('unmute', this.user)
    } else {
      this.$emit('mute', this.user)
    }
  }

  visualize (id, analyser) {
    const canvas:any = this.$refs.canvas
    canvas.width = canvas.width * window.devicePixelRatio
    canvas.height = canvas.height * window.devicePixelRatio
    const canvasCtx = canvas.getContext('2d')
    const bufferLength = analyser.frequencyBinCount
    const dataArray = new Float32Array(bufferLength)
    // const dataArray = new Uint8Array(bufferLength)
    const gb = uuidToColor(id)
    const g = gb[0]
    const b = gb[1]
    const MIN = 7

    function draw () {
      const WIDTH = canvas.width
      const HEIGHT = canvas.height

      analyser.getFloatFrequencyData(dataArray)
      // analyser.getByteTimeDomainData(dataArray)
      // console.log(dataArray[0])
      canvasCtx.fillStyle = 'rgb(0, 0, 0)'
      canvasCtx.fillRect(0, 0, WIDTH, HEIGHT)

      const barWidth = (WIDTH / bufferLength) * 3
      let barHeight = 0
      let point = 0
      let x = 0

      for (let i = 0; i < bufferLength; i++) {
        point = dataArray[i]
        barHeight = (point + 140) * 2

        const r = Math.floor(barHeight + 64)
        if (g % 3 === 0) {
          canvasCtx.fillStyle = `rgb(${r},${g},${b})`
        } else if (g % 3 === 1) {
          canvasCtx.fillStyle = `rgb(${g},${r},${b})`
        } else {
          canvasCtx.fillStyle = `rgb(${g},${b},${r})`
        }

        barHeight = HEIGHT / MIN + barHeight / 256 * HEIGHT * (MIN - 1) / MIN
        if (barHeight < HEIGHT / MIN) {
          barHeight = HEIGHT / MIN
        }
        canvasCtx.fillRect(x, HEIGHT - barHeight, barWidth, barHeight)

        x += barWidth + 2
      }

      setTimeout(function () {
        requestAnimationFrame(draw)
      }, 50)

      // let el = document.getElementById('peer-' + id)
      // if (el && el.getAttribute('data-track-id') === tid) {
      //   setTimeout(function () {
      //     requestAnimationFrame(draw);
      //   }, 50)
      // }
    }

    draw()
  }
}
export default FLoading
</script>

<style lang="scss" scoped>
.user-card {
  border-radius: 20px;
  width: 100%;
  background: rgba(0, 0, 0, 0.3);
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.3);
  position: relative;
  overflow: hidden;
  .top {
    justify-content: space-between;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    .nickname {
      word-break: break-all;
    }
  }
  .canvas-wrapper {
    height: 100px;
    .canvas {
      width: 100%;
      height: 100%;
      border-radius: 20px;
    }
  }
}
</style>
