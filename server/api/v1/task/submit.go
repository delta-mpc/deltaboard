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
	"deltaboard-server/config"
	"deltaboard-server/config/db"
	"deltaboard-server/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Submit(ctx *gin.Context) error {
	node_addr := config.GetConfig().Delta_Node_Addr
	posturl := node_addr + "/v1/task"
	fmt.Println(posturl)
	fmt.Println(ctx.Request.Header.Get("Content-Type"))
	client := &http.Client{Timeout: time.Duration(20) * time.Second}

	reqBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", posturl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	for k := range ctx.Request.Header {
		req.Header.Set(k, ctx.Request.Header.Get(k))
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("resp_getted")
	respTask := &Task{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(respTask); err != nil {
		return err
	}
	auth_token := ctx.Param("token")

	user := &models.User{
		DeltaToken: auth_token,
	}
	if err := db.GetDB().First(user, user).Error; err != nil {
		return err
	}
	task := &models.Task{
		UserId:     user.Id,
		Creator:    user.Name,
		Status:     "SUBMITTED",
		NodeTaskId: respTask.Task_id,
	}
	if err := task.Create(task, db.GetDB()); err != nil {
		return err
	}
	ctx.Writer.Write(body)
	return nil
}
