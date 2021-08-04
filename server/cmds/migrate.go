package cmds

import (
	"deltaboard-server/config"
	"errors"

	"deltaboard-server/config/db"
	"deltaboard-server/config/log"
	"deltaboard-server/migrate"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var migrationInitialized bool

func InitMigration(dbi *gorm.DB) {

	if migrationInitialized {
		return
	}

	// Migration table uses VARCHAR(255) field
	// which cannot be indexed in MySQL before 5.7
	// so we build the migration table using VARCHAR(150)
	// here before the table initialization inside gorm-migrate

	type Migration struct {
		Id string `gorm:"type:varchar(150);primary_key"`
	}

	if err := dbi.AutoMigrate(&Migration{}); err != nil {
		panic(err)
	}

	// Register the actual migrations functions below
	migrate.RegisterMigrations()

	migrationInitialized = true
}

func Migrate(dbi *gorm.DB) error {

	if !migrationInitialized {
		return errors.New("migration not initialized")
	}

	return migrate.Migrate(dbi)
}

func Rollback(dbi *gorm.DB) error {

	if !migrationInitialized {
		return errors.New("migration not initialized")
	}

	return migrate.Rollback(dbi)
}

var _action = "migrate"

func initMigrateCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringVar(&_action, "action", _action, "数据库操作方法(migrate,rollback)")
	return cmd
}

var MigrateCmd = initMigrateCmd(&cobra.Command{
	Use:     "migrate",
	Aliases: []string{"m"},
	Short:   "migrate",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info("migrate begin")
		// Initialize database
		appConfig := config.GetConfig()
		if err := db.InitDB(appConfig); err != nil {
			return err
		}

		InitMigration(db.GetDB())

		switch _action {
		case "migrate":
			err := Migrate(db.GetDB())
			if err != nil {
				log.Errorf("Migrate failed, error:%v", err)
				return err
			}

			log.Info("Migrate succeed!")

			return nil
		case "rollback":
			err := Rollback(db.GetDB())
			if err != nil {
				log.Errorf("Rollback failed, error:%v", err)
				return err
			}

			log.Info("Rollback succeed!")
		default:
			err := errors.New("error action")
			log.Errorf("error action:%v", err)
			return err
		}

		return nil
	},
})
