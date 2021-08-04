package migrate

import (
	"deltaboard-server/migrate/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrationList []*gormigrate.Migration

func Migrate(dbi *gorm.DB) error {
	m := gormigrate.New(dbi, gormigrate.DefaultOptions, migrationList)
	return m.Migrate()
}

func Rollback(dbi *gorm.DB) error {
	m := gormigrate.New(dbi, gormigrate.DefaultOptions, migrationList)
	return m.RollbackLast()
}

func RegisterMigrations() {
	migrationList = append(migrationList, migrations.M202108031056)
}
