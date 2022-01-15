export default function ({ app }) {
  let locale = 'en'
  const browserLang = navigator.language
  if (browserLang.includes('zh')) {
    if (browserLang.includes('zh-TW') || browserLang.includes('zh-HK')) {
      locale = 'zhTW'
    } else {
      locale = 'zh'
    }
  } else if (browserLang.includes('es')) {
    locale = 'es'
  } else if (browserLang.includes('ja')) {
    locale = 'ja'
  } else if (browserLang.includes('de')) {
    locale = 'de'
  }
  app.i18n.setLocale(locale)
}
