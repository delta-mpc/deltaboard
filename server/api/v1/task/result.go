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
	"deltaboard-server/config"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownloadTaskResult(ctx *gin.Context) error {
	node_addr := config.GetConfig().Delta_Node_Addr

	posturl := node_addr + "/v1/task/result?task_id=" + ctx.Param("task_id")
	resp, err := http.Get(posturl)
	if err != nil {
		return err
	}
	for k := range resp.Header {
		ctx.Writer.Header().Set(k, resp.Header.Get(k))
	}
	// respBytes, _ := io.ReadAll(resp.Body)
	// ctx.Writer.Write(respBytes)
	reader := io.TeeReader(resp.Body, ctx.Writer)
	ioutil.ReadAll(reader)
	return nil
}
