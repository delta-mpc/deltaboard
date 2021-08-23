package user

import (
	"crypto/sha256"
	"deltaboard-server/api/v1/response"
	"deltaboard-server/config/db"
	"deltaboard-server/internal/service/current_user"
	"deltaboard-server/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetUserResponse struct {
	response.Response
	Data struct {
		*models.User
		EnterpriseCertificate    string `json:"enterprise_certificate" description:"企业证书"`
		EnterpriseCertificateDNA string `json:"enterprise_certificate_dna" description:"企业证书DNA"`
		Token                    string `json:"user_token" description:"JupyterHub Token"`
	} `json:"data"`
}

type GetUsersInput struct {
	PageID   int `query:"page_id" form:"page_id" description:"page index,begin 0" default:"0"`
	PageSize int `query:"page_size" form:"page_size" description:"page size,default 10" default:"10"`
}
type GetUsersData struct {
	Total int64          `json:"total"`
	List  []*models.User `json:"list"`
}
type GetUserLstResponse struct {
	response.Response
	Data *GetUsersData `json:"data"`
}

func GetUser(ctx *gin.Context) (*GetUserResponse, error) {

	currentUserInterface, _ := ctx.Get("CurrentUser")
	currentUser := currentUserInterface.(*current_user.CurrentUser)

	res := &GetUserResponse{}
	res.Data.User = currentUser.User
	res.Data.Token = fmt.Sprintf("%x", sha256.Sum256([]byte(currentUser.User.Password+currentUser.User.Salt)))
	return res, nil
}

func GetUserList(ctx *gin.Context, in *GetUsersInput) (*GetUserLstResponse, error) {
	var total int64
	if err := db.GetDB().Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, err
	}
	fmt.Println("total ", total)
	users := make([]*models.User, 0)
	if err := db.GetDB().Model(&models.User{}).Offset((in.PageID - 1) * in.PageSize).Limit(in.PageSize).
		Find(&users).Error; err != nil {
		return nil, err
	}
	userData := &GetUsersData{
		Total: total,
		List:  users,
	}
	return &GetUserLstResponse{
		Data: userData,
	}, nil
}
