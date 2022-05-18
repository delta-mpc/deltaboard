
import v1 from './v1';

export default {
   getUserTasks(userId,page,page_size) {
      return v1.get(`/tasks/usertasks/${userId}`,{params:{page:page,page_size:page_size}})
   },
   getTaskLogs(taskId,start,limit) {
      return v1.get(`/tasks/logs`,{params:{task_id:taskId,start:start,limit:limit}})
   },
   getTaskMeta(taskId) {
      return v1.get(`/tasks/meta/${taskId}`)
   },
   getAllTasks(page,page_size){
      return v1.get(`/tasks/all`,{params:{page:page,page_size:page_size}})
   }
};
