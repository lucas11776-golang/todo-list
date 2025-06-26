package database

import (
	"github.com/lucas11776-golang/http/utils/env"
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/databases/sqlite"
)

// Comment
func Setup() {
	orm.DB.Add("sqlite", sqlite.Connect("file:"+env.Env("DATABASE"))) // SQLite database config...
}
