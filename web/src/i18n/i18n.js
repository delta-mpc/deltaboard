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

import enLocale from 'element-ui/lib/locale/lang/en'
import ElementLocale from 'element-ui/lib/locale'

const i18n = new VueI18n({
    locale: 'en',    // 语言标识
    fallbackLocale: 'en',
    messages: {
        en: Object.assign(resources('en'), enLocale)
    }
})

ElementLocale.i18n((key, value) => i18n.t(key, value))

export default i18n;