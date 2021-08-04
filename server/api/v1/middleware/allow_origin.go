package middleware

import (
	"deltaboard-server/config"
	"deltaboard-server/config/origin"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetResponseHeader() gin.HandlerFunc {
	return func(cg *gin.Context) {
		reqOrigin := cg.Request.Header.Get("Origin")

		if reqOrigin != "" && origin.GetAllowOrigin()[reqOrigin] {
			cg.Header("Access-Control-Allow-Origin", reqOrigin)
		}

		if reqOrigin != "" && config.GetConfig().Environment != config.EnvProduction {
			cg.Header("Access-Control-Allow-Origin", reqOrigin)
		}

		cg.Header("Connection", "keep-alive")
		cg.Header("Access-Control-Allow-Credentials", "true")
		cg.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		cg.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method //请求方法
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusNoContent, "Options Request!")
			c.Abort()
		} else {
			c.Next() //  处理请求
		}
	}
}
