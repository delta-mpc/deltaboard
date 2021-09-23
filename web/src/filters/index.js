
import format from 'date-fns/format'
import {zhCN} from 'date-fns/locale'
const filters = {
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
