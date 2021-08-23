
import v1 from './v1';

export default {
   getUserTasks(userId) {
      return v1.get(`/tasks/usertasks/${userId}`)
   },
   getTaskLogs(taskId,page,page_size) {
      return v1.get(`/tasks/logs`,{params:{task_id:taskId,page:page,page_size:page_size}})
   },
   getTaskMeta(taskId) {
      return v1.get(`/tasks/meta/${taskId}`)
   },
   getAllTasks(page,page_size){
      return v1.get(`/tasks/all`,{params:{page:page,page_size:page_size}})
   }
};
