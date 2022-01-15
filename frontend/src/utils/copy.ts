export default {
  onCopy (vue) {
    vue.$toast({ message: vue.$t('common.copy_succ_hint'), color: 'info' })
  },
  onError (_) {
  }
}
