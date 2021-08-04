package v1

import (
	"deltaboard-server/api/v1/middleware"
	"deltaboard-server/api/v1/response"
	"deltaboard-server/api/v1/user"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

func InitRoutes(r *fizz.Fizz) {

	v1g := r.Group("v1", "ApiV1", "API version 1")

	// user group
	UserGroup(v1g)
}

func UserGroup(g *fizz.RouterGroup) {

	userGroup := g.Group("users", "user", "Account APIs")

	userGroup.POST("/tokens", []fizz.OperationOption{
		fizz.Summary("用户注册"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(user.Register, 200))

	userGroup.GET("/auth", []fizz.OperationOption{
		fizz.Summary("用户认证"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(user.UserAuth, 200))

	userGroup.POST("", []fizz.OperationOption{
		fizz.Summary("用户登录"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(user.Login, 200))

	userGroup.GET("", []fizz.OperationOption{
		fizz.Summary("用户信息获取"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, middleware.LoginAuth, tonic.Handler(user.GetUser, 200))

	userGroup.DELETE("/tokens", []fizz.OperationOption{
		fizz.Summary("用户登出"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, middleware.LoginAuth, tonic.Handler(user.Logout, 200))
}
