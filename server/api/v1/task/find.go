/*
 * Copyright 2021 Seven Seals Technology
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package task

import (
	"bytes"
	"deltaboard-server/api/v1/response"
	"deltaboard-server/config"
	"deltaboard-server/config/db"
	"deltaboard-server/config/log"
	"deltaboard-server/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Task_id int64 `json:"task_id"`
}
type UserTask struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Creator      string `json:"creator"`
	Status       string `json:"status"`
	CreatedAt    int64  `json:"created_at"`
	CreatorName  string `json:"creator_name"`
	EnableVerify bool   `json:"enable_verify"`
}

type AllTasks struct {
	Tasks     []*UserTask    `json:"tasks"`
	UserTasks []*models.Task `json:"user_tasks"`
	Total     int            `json:"total_pages"`
}

type GetAllTaskInput struct {
	Page      int64 `json:"page" form:"page" description:"The task page dd"`
	Page_size int64 `json:"page_size" form:"page_size" description:"The task page size"`
}

type GetUserTaskInput struct {
	GetAllTaskInput
	UserId string `json:"userId" form:"userId" path:"userId" description:"The task page size"`
}

type FindTaskResponse struct {
	response.Response
	Data *AllTasks `json:"data"`
}

func FindUserTasks(ctx *gin.Context, in *GetUserTaskInput) (*FindTaskResponse, error) {
	username := in.UserId
	user := models.User{}
	if err := db.GetDB().Select("id", "name").Where("id = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	var total int64
	if err := db.GetDB().Debug().Model(&models.Task{}).Where(" user_id = ?", username).Count(&total).Error; err != nil {
		return nil, err
	}
	tasks := make([]*models.Task, 0)
	if err := db.GetDB().Model(&models.Task{}).Where(" user_id = ?", username).Offset((int(in.Page) - 1) * int(in.Page_size)).Limit(int(in.Page_size)).
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	node_addr := config.GetConfig().Delta_Node_Addr
	posturl := node_addr + "/v1/task/list?task_ids="
	if len(tasks) == 0 {
		return &FindTaskResponse{Data: &AllTasks{
			Tasks: make([]*UserTask, 0),
			Total: 0,
		}}, nil
	}
	for idx, task := range tasks {
		posturl += fmt.Sprintf("%d", task.NodeTaskId)
		if idx < len(tasks)-1 {
			posturl += "&task_ids="
		}
	}
	log.Info(posturl)
	resp, err := http.Get(posturl)
	if err != nil {
		return nil, err
	}
	respBytes, _ := io.ReadAll(resp.Body)
	log.Info(string(respBytes))
	respTasks := make([]*UserTask, 0)
	err = json.NewDecoder(bytes.NewReader(respBytes)).Decode(&respTasks)
	if err != nil {
		return nil, err
	}
	for _, task := range respTasks {
		task.CreatorName = user.Name
	}
	AllTasks := &AllTasks{
		Tasks: respTasks,
		Total: int(total),
	}
	return &FindTaskResponse{Data: AllTasks}, nil
}

func FindAllTasks(ctx *gin.Context, in *GetAllTaskInput) (*FindTaskResponse, error) {
	node_addr := config.GetConfig().Delta_Node_Addr
	page := fmt.Sprintf("%d", in.Page)
	page_size := fmt.Sprintf("%d", in.Page_size)
	posturl := node_addr + "/v1/tasks?page=" + page + "&page_size=" + page_size
	resp, err := http.Get(posturl)
	if err != nil {
		return nil, err
	}
	respTasks := &AllTasks{}
	err = json.NewDecoder(resp.Body).Decode(&respTasks)
	if err != nil {
		return nil, err
	}
	respTaskIds := make([]int64, 0)
	for _, task := range respTasks.Tasks {
		respTaskIds = append(respTaskIds, task.ID)
	}
	userTasks := make([]*models.Task, 0)
	if err := db.GetDB().Where(" node_task_id IN ?", respTaskIds).Find(&userTasks).Error; err != nil {
		return nil, err
	}
	respTasks.UserTasks = userTasks
	return &FindTaskResponse{Data: respTasks}, nil
}
