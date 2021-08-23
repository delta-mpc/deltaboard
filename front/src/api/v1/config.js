
import v1 from './v1';

export default {
   getConfig() {
      return v1.get(`/config`)
   }
};
