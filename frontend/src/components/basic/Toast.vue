<template>
  <v-snackbar v-model="snackbar" :color="bindSnackbar.color" :timeout="3000" top style="top: 60px;">
    {{ bindSnackbar.message }}
  </v-snackbar>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { Mutation, State } from 'vuex-class'

@Component
class Toast extends Vue {
  @State(state => state.app.snackbar) bindSnackbar

  @Mutation('app/setSnackbar') setSnackbar

  get snackbar () {
    return this.bindSnackbar.show
  }

  set snackbar (val) {
    this.setSnackbar(val)
  }
}
export default Toast
</script>
