import LRUCache from 'lru-cache';

const FIVE_MINUTES = 1000 * 60 * 5;
const CAPACITY = 100;
const cache = new LRUCache({ maxAge: FIVE_MINUTES, max: CAPACITY })
export const clearV1Cache = (regex)=>{
   cache.keys().forEach((itm)=>{
      if(regex.test(itm)) {
         cache.del(itm)
      }
   })
}
export default cache