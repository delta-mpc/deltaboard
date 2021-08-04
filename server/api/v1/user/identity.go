package user

import (
	"deltaboard-server/api/v1/response"
	"deltaboard-server/config/db"
	"deltaboard-server/internal/service/current_user"

	"github.com/gin-gonic/gin"
)

type CreateIdentityResponse struct {
	response.Response
	Data *CreateIdentityInput `json:"data"`
}

type CreateIdentityInput struct {
	RealName string `json:"real_name" form:"real_name" validate:"required,gt=1,lt=10" description:"用户真实姓名"`
	CardNo   string `json:"card_no" form:"card_no" validate:"required,gt=17,lt=19" description:"用户身份证号"`
}

func CreateIdentity(ctx *gin.Context, in *CreateIdentityInput) (*CreateIdentityResponse, error) {

	currentUserInterface, _ := ctx.Get("CurrentUser")
	currentUser := currentUserInterface.(*current_user.CurrentUser)

	ok, err := current_user.IdentityCheck(in.RealName, in.CardNo)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, response.NewValidationErrorResponseWithMessage("identity_status", "identity info is invalid")
	}

	currentUser.User.RealName = in.RealName
	currentUser.User.CardNo = in.CardNo

	if err := db.GetDB().Save(currentUser.User).Error; err != nil {
		return nil, err
	}

	return &CreateIdentityResponse{
		Data: in,
	}, nil
}
