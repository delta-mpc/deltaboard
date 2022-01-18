import * as dayjs from "dayjs";
import * as utc from "dayjs/plugin/utc";
import * as timeZone from "dayjs/plugin/timezone";

dayjs.extend(utc);
dayjs.extend(timeZone);

const filters = {
  second2Date: (sec) => {
    return dayjs(sec * 1000).utc().local().format("YYYY-MM-DD HH:mm:ss");
  },
  second2YearMonth: (sec) => {
    return dayjs(sec * 1000).utc().local().format("YYYY-MM")
  },
  second2ChineseYearMonth: (sec) => {
    return dayjs(sec * 1000).utc().tz("Asia/Shanghai").format("YYYY年MM月DD日 E 中国标准时间 a hh:mm:ss")
  },
  privacyPhoneNum: (num) => {
    var reg = /^(\d{3})\d{4}(\d{4})$/;
    return num.replace(reg, "$1****$2");
  },
};
// Object.assign(filters,modules)
export default filters;
