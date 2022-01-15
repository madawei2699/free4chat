<template>
  <v-layout
    v-show="loading || error"
    :class="['f-loading', {'fullscreen': fullscreen }]"
  >
    <v-progress-circular
      v-if="loading"
      :width="3"
      v-bind="$attrs"
      color="primary"
      indeterminate
    />
    <v-btn
      v-if="error"
      icon
      @click="$emit('reload')"
    >
      <v-icon color="primary">
        {{ $icons.mdiReload }}
      </v-icon>
    </v-btn>
  </v-layout>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'

@Component
class FLoading extends Vue {
  @Prop({ type: Boolean, default: false }) fullscreen!: boolean

  @Prop({ type: Boolean, default: false }) loading!: boolean

  @Prop({ type: Boolean, default: false }) error!: boolean
}
export default FLoading
</script>

<style lang="scss" scoped>
.f-loading {
  min-width: 200px;
  padding: 20px;
  display: flex;
  justify-content: center;
  align-items: center;

  &.fullscreen {
    position: fixed;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    z-index: 5;
  }
}
</style>
