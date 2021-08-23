package migrations

import (
	"deltaboard-server/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M202108131837 *gormigrate.Migration

func init() {

	type Task struct {
		models.BaseModel
		UserId  string `json:"userId" gorm:"type:varchar(191);not null"`
		Creator string `json:"name" gorm:"type:varchar(255);not null"`
		Status  string `json:"status" gorm:"type:varchar(16);not null"`
	}

	M202108131837 = &gormigrate.Migration{
		ID: "M202108131837",
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
