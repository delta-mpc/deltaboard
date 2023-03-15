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

type UserList struct {
	List  []*models.User `json:"list"`
	Total int64          `json:"total"`
}

type FetchUserInput struct {
	ApproveStatus int8  `json:"approve_status" form:"approve_status"`
	Page          int64 `json:"page" form:"page"`
	PageSize      int64 `json:"page_size" form:"page_size"`
	Sort          int8  `json:"sort" form:"sort" validate:"gte=0,lte=1" description:"结果的顺序。0-ASC,1-DESC"`
}

type FetchRegistedUserResponse struct {
	response.Response
	Data *UserList `json:"data"`
}

func Register(ctx *gin.Context, in *RegisterInput) (*RegisterResponse, error) {

	ctxuser, exist := ctx.Get("CurrentUser")
	isOperatedByAdmin := exist && ctxuser.(*current_user.CurrentUser).User.Role == models.ROLE_ADMIN
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
	if isOperatedByAdmin {
		user.ApprvStatus = models.USER_APPROV_STATUS_APPROVED
	} else {
		user.ApprvStatus = models.USER_APPROV_STATUS_REGISTED
	}
	if err := user.Create(user, db.GetDB()); err != nil {
		return nil, err
	}
	return &RegisterResponse{
		Data: user,
	}, nil
}

func Approve(ctx *gin.Context) (*response.Response, error) {
	ctxuser, exist := ctx.Get("CurrentUser")
	isOperatedByAdmin := exist && ctxuser.(*current_user.CurrentUser).User.Role == models.ROLE_ADMIN
	if !isOperatedByAdmin {
		return nil, response.NewValidationErrorResponseWithMessage("user_approve", "unsufficient privilege")
	}
	userId := ctx.Param("userId")
	user := &models.User{
		BaseModel: models.BaseModel{
			Id: userId,
		},
	}
	if err := db.GetDB().First(user, user).Error; err != nil {
		return nil, response.NewValidationErrorResponseWithMessage("user_register", "user_not_exist")
	}
	user.ApprvStatus = models.USER_APPROV_STATUS_APPROVED
	db.GetDB().Save(user)
	return &response.Response{Message: "success"}, nil
}

func Reject(ctx *gin.Context) (*response.Response, error) {
	ctxuser, exist := ctx.Get("CurrentUser")
	isOperatedByAdmin := exist && ctxuser.(*current_user.CurrentUser).User.Role == models.ROLE_ADMIN
	if !isOperatedByAdmin {
		return nil, response.NewValidationErrorResponseWithMessage("user_approve", "unsufficient privilege")
	}
	userId := ctx.Param("userId")
	user := &models.User{
		BaseModel: models.BaseModel{
			Id: userId,
		},
	}
	if err := db.GetDB().First(user, user).Error; err != nil {
		return nil, response.NewValidationErrorResponseWithMessage("user_register", "user_not_exist")
	}
	user.ApprvStatus = models.USER_APPROV_STATUS_REJECTED
	db.GetDB().Save(user)
	return &response.Response{Message: "success"}, nil
}

func DelUser(ctx *gin.Context) (*response.Response, error) {
	ctxuser, exists := ctx.Get("CurrentUser")
	if !exists {
		return nil, response.NewValidationErrorResponseWithMessage("user_delete", "unsufficient privilege")
	}
	if ctxuser.(*current_user.CurrentUser).Role != models.ROLE_ADMIN {
		return nil, response.NewValidationErrorResponseWithMessage("user_delete", "unsufficient privilege")
	}
	user := &models.User{
		BaseModel: models.BaseModel{
			Id: ctx.Param("userId"),
		},
	}
	if err := db.GetDB().First(user, user).Error; err != nil {
		return nil, err
	}
	tx := db.GetDB().Begin()
	db.GetDB().Delete(user)
	tx.Commit()
	return &response.Response{Message: "success"}, nil
}

func FetchUser(ctx *gin.Context, in *FetchUserInput) (*FetchRegistedUserResponse, error) {
	var total int64
	if err := db.GetDB().Model(&models.User{}).Where(" apprv_status = ? ", in.ApproveStatus).Count(&total).Error; err != nil {
		return nil, err
	}
	users := make([]*models.User, 0)
	if total > 0 {
		var order string
		if in.Sort == 0 {
			order = "created_at asc"
		} else {
			order = "created_at desc"
		}
		if err := db.GetDB().Model(&models.User{}).Where("apprv_status = ? ", in.ApproveStatus).Order(order).
			Limit(int(in.PageSize)).Offset(int((in.Page - 1) * in.PageSize)).Find(&users).Error; err != nil {
			return nil, err
		}
	}
	return &FetchRegistedUserResponse{
		Data: &UserList{
			List:  users,
			Total: total,
		},
	}, nil
}
