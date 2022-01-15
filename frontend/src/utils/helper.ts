export function errorHandler (vue: Vue, error: { message: string; code: string | number }) {
  const $toast = vue.$utils.toast
  const fallback = '未知错误'
  const message = `${error.code || ''} ${error.message || fallback}`
  $toast(vue, { message, color: 'error' })
}

export function toast (vue: Vue, data: { message: string; color?: string }) {
  vue.$store.commit('app/toast', data)
}

export function reloadMixinTheme () {
  const w: any = window
  let env = ''
  if (
    w.webkit &&
    w.webkit.messageHandlers &&
    w.webkit.messageHandlers.MixinContext
  ) {
    env = 'iOS'
  }
  if (w.MixinContext && w.MixinContext.getContext) {
    env = 'Android'
  }
  switch (env) {
    case 'iOS':
      w.webkit.messageHandlers.reloadTheme.postMessage('')
      return
    case 'Android':
      w.MixinContext.reloadTheme()
  }
}
