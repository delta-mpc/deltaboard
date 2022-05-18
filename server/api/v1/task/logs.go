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
	"deltaboard-server/api/v1/response"
	"deltaboard-server/config"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskLog struct {
	ID        int64  `json:"id"`
	CreatedAt int64  `json:"created_at"`
	Message   string `json:"message"`
	TxHash    string `json:"tx_hash"`
}

type TaskLogInput struct {
	TaskId int64 `json:"task_id" form:"task_id" validate:"required" description:"The task id"`
	Start  int64 `json:"start" form:"start" description:"The task log start id"`
	Limit  int64 `json:"limit" form:"limit" description:"The task log query limit size"`
}

type TaskLogResponse struct {
	response.Response
	Data []*TaskLog `json:"data"`
}

func GetTaskLogs(ctx *gin.Context, in *TaskLogInput) (*TaskLogResponse, error) {
	node_addr := config.GetConfig().Delta_Node_Addr
	taskId := fmt.Sprintf("%d", in.TaskId)
	start := fmt.Sprintf("%d", in.Start)
	limit := fmt.Sprintf("%d", in.Limit)
	posturl := node_addr + "/v1/task/logs?task_id=" + taskId + "&start=" + start + "&limit=" + limit
	client := &http.Client{Timeout: time.Duration(20) * time.Second}
	resp, err := client.Get(posturl)
	if err != nil {
		return nil, err
	}
	taskLogs := make([]*TaskLog, 0)
	if err := json.NewDecoder(resp.Body).Decode(&taskLogs); err != nil {
		return nil, err
	}
	return &TaskLogResponse{
		Data: taskLogs,
	}, nil
}
