import Vue from 'vue'
import {
  mdiMicrophone,
  mdiMicrophoneOff,
  mdiAccountCircleOutline,
  mdiContentCopy,
  mdiClose,
  mdiRefresh,
  mdiDice3,
  mdiHome,
  mdiHeadphones,
  mdiArrowLeft
} from '@mdi/js'
import * as utils from '@/utils'

Vue.prototype.$icons = {
  mdiMicrophone,
  mdiMicrophoneOff,
  mdiAccountCircleOutline,
  mdiContentCopy,
  mdiClose,
  mdiRefresh,
  mdiDice3,
  mdiHome,
  mdiHeadphones,
  mdiArrowLeft
}

Vue.prototype.$utils = utils

Vue.prototype.$toast = function (data: { message: string; color: string }) {
  this.$store.commit('app/toast', data)
}
