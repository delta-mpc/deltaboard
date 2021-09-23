
import v1 from './v1';

export default {
   listNodes(page,page_size) {
      return v1.get(`/nodes`,{params:{page:page,page_size:page_size}})
   }
};
