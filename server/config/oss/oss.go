package oss

import (
	"deltaboard-server/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var conn *oss.Client

func InitOss(appConfig *config.AppConfig) (err error) {

	conn, err = oss.New(appConfig.Oss.Endpoint, appConfig.Oss.AccessKeyID, appConfig.Oss.AccessKeySecret)
	return err
}

func GetOss() *oss.Client {
	return conn
}
