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
	"deltaboard-server/api/v1/response"
	"deltaboard-server/config/db"
	"deltaboard-server/internal/service/current_user"
	"deltaboard-server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	UserValidationKey = "validation_code"
	PhoneCode         = "phone_code"
)

type RegisterInput struct {
	UserName    string `json:"user_name" form:"user_name" validate:"required" description:"The user's name"`
	UserEmail   string `json:"user_email" form:"user_email" description:"The user's email"`
	Password    string `json:"password" form:"password" validate:"required" description:"The user's encrypted password"`
	PhoneNumber string `json:"phonenumber" form:"phonenumber" description:"The user's phonenumber"`
}

type GenerateCodeInput struct {
	PhoneNumber string `json:"phonenumber" form:"phonenumber"`
}

type RegisterResponse struct {
	response.Response
	Data *models.User `json:"data"`
}

func Register(ctx *gin.Context, in *RegisterInput) (*RegisterResponse, error) {

	user := &models.User{
		Name: in.UserName,
	}
	//查重
	if err := db.GetDB().First(user, user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	} else {
		return nil, response.NewValidationErrorResponseWithMessage("user_register", "user_already_exist")
	}

	user = models.NewUser(in.UserName, in.Password, in.UserEmail, in.PhoneNumber)
	if err := user.Create(user, db.GetDB()); err != nil {
		return nil, err
	}

	//设置cookie
	currentUser := &current_user.CurrentUser{
		User: user,
		UserVerification: current_user.UserVerification{
			UserId: user.Id,
		},
	}
	if err := currentUser.SaveUserLoginSession(ctx.Request, ctx.Writer); err != nil {
		return nil, err
	}

	return &RegisterResponse{
		Data: user,
	}, nil
}
