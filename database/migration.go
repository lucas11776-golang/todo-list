package database

import (
	models "server/database/migrations"

	"github.com/lucas11776-golang/http/utils/env"
	"github.com/lucas11776-golang/orm/migrations"
)

// Comment
func Migrate(envPath string) {
	env.Load(".env")
	Setup()

	migrations.Migrations(
		&models.UserMigration{},
		&models.TaskMigration{},
	).Up()
}
