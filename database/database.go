package database

import (
	"github.com/lucas11776-golang/http/utils/env"
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/drivers/sqlite"
)

// Comment
func Setup() orm.Database {
	db := sqlite.Connect(env.Env("DATABASE"))

	orm.DB.Add(orm.DefaultDatabaseName, db) // SQLite database config...

	return db
}
