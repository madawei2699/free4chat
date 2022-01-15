
import * as utils from '@/utils'

declare module 'vue/types/vue' {
  interface Vue {
    title?: string
    $errorHandler: (vue: Vue, error: { message: string, code: string | number }) => void
    $toast: (vue: Vue, data: { message: string, color: string }) => void
    $utils: typeof utils
  }
}

export * from './enum'
