package v1

import (
	app_config "deltaboard-server/api/v1/config"
	node "deltaboard-server/api/v1/delta_node"
	"deltaboard-server/api/v1/middleware"
	"deltaboard-server/api/v1/response"
	"deltaboard-server/api/v1/task"
	"deltaboard-server/api/v1/user"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

func InitRoutes(r *fizz.Fizz) {

	v1g := r.Group("v1", "ApiV1", "API version 1")
	// user group
	UserGroup(v1g)

	TaskGroup(v1g)

	ConfigGroup(v1g)

	NodeGroup(v1g)

	r.POST("/submit/:token/v1/task", []fizz.OperationOption{
		fizz.Summary("上传任务"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(task.Submit, 200))
}

func NodeGroup(g *fizz.RouterGroup) {
	nodeGroup := g.Group("nodes", "node", "Node APIs")
	nodeGroup.GET("", []fizz.OperationOption{
		fizz.Summary("获取节点"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(node.GetNodes, 200))
}

func TaskGroup(g *fizz.RouterGroup) {
	taskGroup := g.Group("tasks", "task", "Account APIs")
	// taskGroup.POST("/:token/v1/task", []fizz.OperationOption{
	// 	fizz.Summary("上传任务"),
	// 	fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	// }, tonic.Handler(task.Submit, 200))

	taskGroup.GET("/meta/:taskId", []fizz.OperationOption{
		fizz.Summary("上传任务"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(task.GetTaskMetaData, 200))

	taskGroup.GET("/logs", []fizz.OperationOption{
		fizz.Summary("上传任务"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(task.GetTaskLogs, 200))

	taskGroup.GET("/all", []fizz.OperationOption{
		fizz.Summary("所有任务获取"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(task.FindAllTasks, 200))

	taskGroup.GET("/usertasks/:userId", []fizz.OperationOption{
		fizz.Summary("用户任务获取"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(task.FindUserTasks, 200))

	taskGroup.GET("/result/:task_id", []fizz.OperationOption{
		fizz.Summary("下载结果"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(task.DownloadTaskResult, 200))
}

func ConfigGroup(g *fizz.RouterGroup) {
	configGroup := g.Group("config", "config", "App Configs")
	configGroup.GET("", []fizz.OperationOption{
		fizz.Summary("获取应用配置"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(app_config.Config, 200))
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

	userGroup.GET("/list", []fizz.OperationOption{
		fizz.Summary("用户列表"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(user.GetUserList, 200))

	userGroup.GET("", []fizz.OperationOption{
		fizz.Summary("用户信息获取"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, middleware.LoginAuth, tonic.Handler(user.GetUser, 200))

	userGroup.PUT("/password", []fizz.OperationOption{
		fizz.Summary("用户重置密码"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, middleware.LoginAuth, tonic.Handler(user.ResetLoginPassword, 200))

	userGroup.DELETE("/tokens", []fizz.OperationOption{
		fizz.Summary("用户登出"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, middleware.LoginAuth, tonic.Handler(user.Logout, 200))

	userGroup.DELETE("/del/:userId", []fizz.OperationOption{
		fizz.Summary("删除用户"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, middleware.LoginAuth, tonic.Handler(user.DelUser, 200))

	userGroup.POST("/delta_token/renew", []fizz.OperationOption{
		fizz.Summary("刷新链接"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, middleware.LoginAuth, tonic.Handler(user.RenewDeltaToken, 200))

	userGroup.POST("/approve/:userId", []fizz.OperationOption{
		fizz.Summary("审核通过"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, middleware.LoginAuth, tonic.Handler(user.Approve, 200))

	userGroup.POST("/reject/:userId", []fizz.OperationOption{
		fizz.Summary("审核不通过"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, middleware.LoginAuth, tonic.Handler(user.Reject, 200))

	userGroup.GET("/approve_status", []fizz.OperationOption{
		fizz.Summary("获取审核状态"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, middleware.LoginAuth, tonic.Handler(user.FetchUser, 200))

}
