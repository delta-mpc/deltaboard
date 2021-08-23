package user

import (
	"deltaboard-server/api/v1/response"
	"deltaboard-server/config/db"
	"deltaboard-server/internal/library/random"
	"deltaboard-server/internal/service/current_user"
	"deltaboard-server/models"

	"github.com/gin-gonic/gin"
)

type NewDeltaTokenResponse struct {
	response.Response
	Data *struct {
		Token string `json:"new_token"`
	} `json:"data"`
}

func RenewDeltaToken(ctx *gin.Context) (*NewDeltaTokenResponse, error) {
	ctxuser, exists := ctx.Get("CurrentUser")
	if !exists {
		return &NewDeltaTokenResponse{
			Response: response.Response{
				Message: "用户未登录",
			},
		}, nil
	}
	user := &models.User{
		BaseModel: models.BaseModel{
			Id: ctxuser.(*current_user.CurrentUser).User.Id,
		},
	}
	if err := db.GetDB().Find(user, user).Error; err != nil {
		return nil, err
	}
	user.DeltaToken = random.GenerateRandomString(32)
	db.GetDB().Save(user)
	return &NewDeltaTokenResponse{
		Data: &struct {
			Token string `json:"new_token"`
		}{
			Token: user.DeltaToken,
		},
	}, nil
}
