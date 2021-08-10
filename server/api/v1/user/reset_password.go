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

package user

import (
	"crypto/sha256"
	"deltaboard-server/api/v1/response"
	"deltaboard-server/config/db"
	"deltaboard-server/internal/service/current_user"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ResetLoginPasswordInput struct {
	OldPassword string `json:"old_password" form:"old_password" validate:"required,gt=4,lt=127" description:"旧密码"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required,gt=4,lt=127" description:"新密码"`
}

func ResetLoginPassword(ctx *gin.Context, in *ResetLoginPasswordInput) (*response.Response, error) {

	currentUserInterface, _ := ctx.Get("CurrentUser")
	currentUser := currentUserInterface.(*current_user.CurrentUser)

	oldPwdHash := fmt.Sprintf("%x", sha256.Sum256([]byte(in.OldPassword)))
	if currentUser.Password != fmt.Sprintf("%x", sha256.Sum256([]byte(oldPwdHash+currentUser.Salt))) {
		return nil, response.NewValidationErrorResponseWithMessage("old_password", "invalid_password")
	}

	newPwdHash := fmt.Sprintf("%x", sha256.Sum256([]byte(in.NewPassword)))
	newPwdHash = fmt.Sprintf("%x", sha256.Sum256([]byte(newPwdHash+currentUser.Salt)))
	//更新
	currentUser.Password = newPwdHash
	if err := db.GetDB().Save(currentUser.User).Error; err != nil {
		return nil, err
	}
	return &response.Response{}, currentUser.RemoveUserLoginSession(ctx.Request, ctx.Writer)
}
