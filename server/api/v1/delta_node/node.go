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
package node

import (
	"deltaboard-server/api/v1/request"
	"deltaboard-server/api/v1/response"
	"deltaboard-server/config"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeltaNode struct {
	ID   string `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

type DeltaNodeList struct {
	List  []*DeltaNode `json:"nodes"`
	Total int          `json:"total_pages"`
}

type ListNodeResponse struct {
	response.Response
	Data *struct {
		List  []*DeltaNode `json:"list"`
		Total int          `json:"total"`
	} `json:"data"`
}

func GetNodes(ctx *gin.Context, in *request.PageInput) (*ListNodeResponse, error) {
	node_addr := config.GetConfig().Delta_Node_Addr
	page := fmt.Sprintf("%d", in.Page)
	page_size := fmt.Sprintf("%d", in.PageSize)
	posturl := node_addr + "/v1/nodes?page=" + page + "&page_size=" + page_size
	resp, err := http.Get(posturl)
	if err != nil {
		return nil, err
	}
	deltaNodes := &DeltaNodeList{}
	if err := json.NewDecoder(resp.Body).Decode(&deltaNodes); err != nil {
		return nil, err
	}
	return &ListNodeResponse{
		Data: &struct {
			List  []*DeltaNode `json:"list"`
			Total int          `json:"total"`
		}{
			List:  deltaNodes.List,
			Total: deltaNodes.Total,
		},
	}, nil
}
