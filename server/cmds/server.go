package cmds

import (
	"deltaboard-server/config/db"
	"deltaboard-server/config/origin"
	"deltaboard-server/config/session"
	"fmt"

	"deltaboard-server/api"
	"deltaboard-server/config"
	"deltaboard-server/config/log"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	RunE: func(cmd *cobra.Command, args []string) error {

		log.Info("start service server")
		appConfig := config.GetConfig()

		// Initialize app config
		if err := InitAllFromAppConfig(appConfig); err != nil {
			panic(err)
		}

		// Start http server
		app := api.GetHttpApplication(appConfig)
		address := fmt.Sprintf("%s:%s", appConfig.Http.Host, appConfig.Http.Port)

		log.Info("server url:" + address)
		return app.Run(address)
	},
}

func InitAllFromAppConfig(appConfig *config.AppConfig) error {

	// Initialize database
	if err := db.InitDB(appConfig); err != nil {
		return err
	}

	// Initialize session store
	if err := session.InitSessionStore(appConfig); err != nil {
		return err
	}

	// Initialize cors allow origin
	if err := origin.InitOrigin(appConfig); err != nil {
		return err
	}

	return nil
}
