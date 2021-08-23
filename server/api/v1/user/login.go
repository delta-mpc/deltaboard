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
	"deltaboard-server/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	UserName string `json:"user_name" form:"user_name" validate:"required" description:"The user's Name"`
	Password string `json:"password" form:"password" validate:"required,gt=4,lt=127" description:"The user's encrypted password"`
}

type AuthInput struct {
	UserName string `json:"user_name" form:"user_name" validate:"required" description:"The user's name"`
	Token    string `json:"token" form:"token" validate:"required" description:"The user's encrypted password"`
}

type LoginResponse struct {
	response.Response
	Data models.User `json:"data"`
}

func UserAuth(ctx *gin.Context, in *AuthInput) (*LoginResponse, error) {
	user := &models.User{
		Name: in.UserName,
	}
	//查重
	if err := db.GetDB().First(user, user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, response.NewValidationErrorResponseWithMessage("user_login", "user_not_found")
		}
		return nil, err
	}

	if user.ApprvStatus != models.USER_APPROV_STATUS_APPROVED {
		return nil, response.NewValidationErrorResponseWithMessage("user_login", "user_not_approved")
	}

	if in.Token != fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password+user.Salt))) {
		return nil, response.NewValidationErrorResponseWithMessage("user_auth", "invalid_token")
	}
	return &LoginResponse{
		Data: *user,
	}, nil
}

func Login(ctx *gin.Context, in *LoginInput) (*LoginResponse, error) {
	fmt.Println("user_name:" + in.UserName)
	user := &models.User{
		Name: in.UserName,
	}

	//查重
	if err := db.GetDB().First(user, user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, response.NewValidationErrorResponseWithMessage("user_login", "user_not_found")
		}
		return nil, err
	}

	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(in.Password)))

	if user.Password != fmt.Sprintf("%x", sha256.Sum256([]byte(passwordHash+user.Salt))) {
		return nil, response.NewValidationErrorResponseWithMessage("user_login", "invalid_password")
	}

	// 更新用户登录状态至session
	currentUser := &current_user.CurrentUser{
		User: user,
		UserVerification: current_user.UserVerification{
			UserId: user.Id,
		},
	}

	if err := currentUser.SaveUserLoginSession(ctx.Request, ctx.Writer); err != nil {
		return nil, err
	}

	return &LoginResponse{
		Data: *user,
	}, nil
}
