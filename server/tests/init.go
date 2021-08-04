package tests

import (
	"deltaboard-server/api"
	"deltaboard-server/cmds"
	"deltaboard-server/config"
	"deltaboard-server/config/log"

	"github.com/gin-gonic/gin"
)

var Application *gin.Engine = nil

func init() {
	// Initialize api application to serve API test calls

	testAppConfig := &config.AppConfig{}

	testAppConfig.Environment = config.EnvTest
	testAppConfig.Db.Driver = ""
	testAppConfig.Log.Level = log.StandardLogger().Level.String()
	testAppConfig.Http.Host = "127.0.0.1"
	testAppConfig.Http.Port = "8080"

	if err := cmds.InitAllFromAppConfig(testAppConfig); err != nil {
		panic(err)
	}

	Application = api.GetHttpApplication(testAppConfig)

}
