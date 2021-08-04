package migrations

import (
	"deltaboard-server/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M202108031056 *gormigrate.Migration

func init() {

	type User struct {
		models.BaseModel
		Name     string            `json:"name" gorm:"type:varchar(127);not null"`
		Password string            `json:"-" gorm:"type:varchar(64);not null"`
		Salt     string            `json:"-" gorm:"type:varchar(32);not null"`
		Email    string            `json:"email" gorm:"type:varchar(127);unique;not null"`
		Status   models.UserStatus `json:"status" gorm:"type:tinyint(1);not null"`
		RealName string            `json:"real_name" gorm:"type:varchar(32);not null"`
		CardNo   string            `json:"card_no" gorm:"type:varchar(32);not null"`
		Phone    string            `json:"phonenumber" gorm:"type:varchar(11);not null"`
	}

	M202108031056 = &gormigrate.Migration{
		ID: "M202108031056",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(&User{}); err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable("users"); err != nil {
				return err
			}
			return nil
		},
	}
}
