package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var appConfig *AppConfig

// Init is an exported method that takes the config from the config file
// and unmarshal it into AppConfig struct
func InitConfig(configPath string) error {
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName("config")
	if configPath != "" {
		v.AddConfigPath(configPath)
		// v.AddConfigPath("E:/delta-board/delta-board-server")
	} else {
		// v.AddConfigPath(configPath)
		v.AddConfigPath("config")
	}

	if err := v.ReadInConfig(); err != nil {
		panic("Read config file failed:" + err.Error())
	}

	appConfig = &AppConfig{}

	if err := v.Unmarshal(appConfig); err != nil {
		panic("Parse config file failed:" + err.Error())
	}

	return nil
}

func GetConfig() *AppConfig {
	return appConfig
}
