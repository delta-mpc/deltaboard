package migrations

import (
	"deltaboard-server/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M202108181512 *gormigrate.Migration

func init() {

	type User struct {
		models.BaseModel
		Name        string `json:"name" gorm:"type:varchar(127);not null"`
		Password    string `json:"-" gorm:"type:varchar(64);not null"`
		Salt        string `json:"-" gorm:"type:varchar(32);not null"`
		Email       string `json:"email" gorm:"type:varchar(127)"`
		RealName    string `json:"real_name" gorm:"type:varchar(32)"`
		CardNo      string `json:"card_no" gorm:"type:varchar(32)"`
		Phone       string `json:"phonenumber" gorm:"type:varchar(11)"`
		Role        int64  `json:"role" gorm:"type:tinyint(8)"`
		DeltaToken  string `json:"delta_token" gorm:"type:varchar(191);" description:"auth token"`
		ApprvStatus int8   `json:"approve_status" gorm:"type:tinyint(8)"`
	}

	M202108181512 = &gormigrate.Migration{
		ID: "M202108181512",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(&User{}); err != nil {
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
