package migrations

import (
	"deltaboard-server/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M2021081161553 *gormigrate.Migration

func init() {

	type Task struct {
		models.BaseModel
		UserId     string `json:"userId" gorm:"type:varchar(191)"`
		Creator    string `json:"name" gorm:"type:varchar(255)"`
		Status     string `json:"status" gorm:"type:varchar(16)"`
		NodeTaskId int64  `json:"nodeTaskId" gorm:"type:tinyint(64)"`
	}

	M2021081161553 = &gormigrate.Migration{
		ID: "M2021081161553",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(&Task{}); err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable("tasks"); err != nil {
				return err
			}
			return nil
		},
	}
}
