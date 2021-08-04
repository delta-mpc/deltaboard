package origin

import "deltaboard-server/config"

var AllowOrigin map[string]bool

func InitOrigin(appConfig *config.AppConfig) (err error) {

	AllowOrigin = make(map[string]bool)
	for _, origin := range appConfig.AllowOrigins {
		AllowOrigin[origin] = true
	}

	return nil
}

func GetAllowOrigin() map[string]bool {

	return AllowOrigin
}
