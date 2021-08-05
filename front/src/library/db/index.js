import schema from './schema.js'
const dbName = 'ipproof'
export const DB = {
   db: null,
   open: function () {
      const self = this;
      if (this.isConnected()) {
         return new Promise((resolve,reject)=>{
            resolve(self.db)
         })
      } else {
         return new Promise((resolve, reject) => {
            let idxdb = indexedDB.open(dbName, schema.version);
            idxdb.onsuccess = function (event) {
               self.db = event.target.result;
               resolve(self.db)
            };
            idxdb.onerror = function (event) {
               reject('open indexed db failed')
            };
            idxdb.onupgradeneeded = function (event) {
               var db = event.target.result
               if (schema.stores) {
                  schema.stores.forEach((store) => {
                     if(!db.objectStoreNames.contains(store.name))
                        db.createObjectStore(store.name, store)
                  })
               }
            }
         })
      }
   },
   insert: function (storeName, entity) {
      const self = this;
      return this.open().then(()=>new Promise((resolve, reject) => {
         let req = self.db.transaction(storeName, "readwrite").objectStore(storeName).add(entity);
         req.onsuccess = function (event) {
            resolve(event.target.result)
         };
         req.onerror = function (event) {
            reject()
         }
      }))
   },
   select: function (storeName, key) {
      const self = this;
      return this.open().then(()=>new Promise((resolve, reject) => {
         let transaction = self.db.transaction(storeName);
         let obStore = transaction.objectStore(storeName),
            obReq = obStore.get(key);
         obReq.onerror = function (event) {
            reject()
         }
         obReq.onsuccess = function (event) {
            resolve(event.target.result)
         }
      }))
   },
   del: function (storeName, key) {
      const self = this;
      return this.open().then(()=>new Promise((resolve, reject) => {
         let request = self.db.transaction(storeName, "readwrite")
            .objectStore(storeName)
            .delete(key);
         request.onsuccess = function (event) {
            resolve(self.db)
         };
         request.onerror = function (event) {
            reject()
         };
      }))
   },
   isConnected: function () {
      return this.db != null
   }
}