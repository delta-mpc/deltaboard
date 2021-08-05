import axios from 'axios'
import {DB} from '@/library/db'
let httpClient = axios.create({
});
export const fetchFile = function(baseUrl,filePath) {
   return DB.select('static_files',filePath).then((res)=>{
      if(res) {
         return res.file.arrayBuffer().then((buffer)=>new Uint8Array(buffer));
      } else {
         return httpClient.get(`${baseUrl}${filePath}`,
         {responseType:'blob'}).then((resp)=>{
            return DB.insert('static_files',{
               file_name:filePath,
               file:resp.data
            }).then((record)=>{
               console.log('inserted',filePath);
               return resp.data.arrayBuffer().then((buffer)=>new Uint8Array(buffer));
            })
         })
      }
   })
}