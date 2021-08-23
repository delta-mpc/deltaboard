package task

import (
	"deltaboard-server/config"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskMetaData(ctx *gin.Context) (*FindTaskResponse, error) {
	taskId := ctx.Param("taskId")
	node_addr := config.GetConfig().Delta_Node_Addr
	posturl := node_addr + "/v1/task/metadata?task_id=" + taskId
	resp, err := http.Get(posturl)
	if err != nil {
		return nil, err
	}
	respTasks := &UserTask{}
	if err := json.NewDecoder(resp.Body).Decode(respTasks); err != nil {
		return nil, err
	}
	respTaskLst := make([]*UserTask, 0)
	respTaskLst = append(respTaskLst, respTasks)
	return &FindTaskResponse{Data: &AllTasks{
		Tasks: respTaskLst,
		Total: 1,
	}}, nil
}
