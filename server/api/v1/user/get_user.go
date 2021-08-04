package user

import (
	"crypto/sha256"
	"deltaboard-server/api/v1/response"
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

func GetUser(ctx *gin.Context) (*GetUserResponse, error) {

	currentUserInterface, _ := ctx.Get("CurrentUser")
	currentUser := currentUserInterface.(*current_user.CurrentUser)

	res := &GetUserResponse{}
	res.Data.User = currentUser.User
	res.Data.Token = fmt.Sprintf("%x", sha256.Sum256([]byte(currentUser.User.Password+currentUser.User.Salt)))
	return res, nil
}
