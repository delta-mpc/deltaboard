
import format from 'date-fns/format'
import {zhCN} from 'date-fns/locale'
const statusArr = ['', '待审核', '审核通过', '审核不通过', '存证中', '存证成功', '存证失败', '被更新']

// const modulesFiles = require.context('./modules', true, /\.js$/)

// const modules = modulesFiles.keys().reduce((modules, modulePath) => {
//     const moduleName = modulePath.replace(/^\.\/(.*)\.\w+$/,'$1')
//     const value = modulesFiles(modulePath)
//       Object.keys(value.default).forEach((itm)=>{
//          modules[`${moduleName}.${itm}`] = value.default[itm]
//       })
//       return modules
// }, {})
const filters = {
  fileStatus2Txt: (value) => {
    return statusArr[value]
  },
  second2Date: (sec) => {
    return format(sec * 1000, 'yyyy-MM-dd HH:mm:ss')
  },
  second2YearMonth: (sec) => {
    return format(sec * 1000, 'yyyy-MM')
  },
  second2ChineseYearMonth: (sec) => {
    return format(sec * 1000, 'yyyy年MM月dd日 E 中国标准时间 a hh:mm:ss',{locale:zhCN})
  },
  privacyPhoneNum: (num) => {
    var reg = /^(\d{3})\d{4}(\d{4})$/
    return num.replace(reg, "$1****$2");
  }
}
// Object.assign(filters,modules)
export default filters
