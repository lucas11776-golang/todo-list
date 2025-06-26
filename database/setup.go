package database

import (
	"github.com/lucas11776-golang/http/utils/env"
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/databases/sqlite"
)

// Comment
func Setup() orm.Database {
	db := sqlite.Connect("file:" + env.Env("DATABASE"))

	orm.DB.Add("sqlite", db) // SQLite database config...

	return db
}
