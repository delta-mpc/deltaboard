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
	migrationList = append(migrationList, migrations.M202108131837)
	migrationList = append(migrationList, migrations.M202108132001)
	migrationList = append(migrationList, migrations.M2021081161553)
	migrationList = append(migrationList, migrations.M202108171046)
	migrationList = append(migrationList, migrations.M202108181512)
}
