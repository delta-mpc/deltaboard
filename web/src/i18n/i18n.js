import Vue from 'vue'
import VueI18n from 'vue-i18n'
const resources = (lang) => {
   const requireContext = require.context('@/i18n',true, /.*\/.*\.js$/)
   return  requireContext.keys()
    .filter((itm)=>new RegExp(`/${lang}/`).test(itm))
    .reduce((pre,cur)=>{
      return {
         ...(requireContext(cur).default || requireContext(cur)),
         ...pre
      }
    },{})
}
Vue.use(VueI18n)

import zhLocale from 'element-ui/lib/locale/lang/zh-CN'
import enLocale from 'element-ui/lib/locale/lang/en'
import ElementLocale from 'element-ui/lib/locale'

var language = (navigator.browserLanguage || navigator.language).toLowerCase();
var localeLang = '';
if (language.indexOf('zh') > -1) {
    localeLang = 'zh'
} else {
    localeLang = 'en'
}
const i18n = new VueI18n({
    locale: localeLang,    // 语言标识
    // locale: 'zh',    // 语言标识
    fallbackLocale: 'zh',
    messages: {
        zh: Object.assign(resources('zh'), zhLocale),
        en: Object.assign(resources('en'), enLocale)
    }
})

ElementLocale.i18n((key, value) => i18n.t(key, value))

export default i18n;