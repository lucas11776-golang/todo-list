package migrations

import (
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/migrations"
)

type UserMigration struct{}

// Comment
func (ctx *UserMigration) Up() {
	migrations.Create(orm.DefaultDatabaseName, "users", func(table *migrations.Table) {
		table.Increment("id")
		table.TimeStamp("created").Current()
		table.String("first_name")
		table.String("last_name")
		table.String("email")
		table.String("password")
	})
}

// Comment
func (ctx *UserMigration) Down() {
	migrations.Drop("sqlite", "users")
}
