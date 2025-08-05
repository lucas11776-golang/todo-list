package migrations

import (
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/migrations"
)

type TaskMigration struct{}

// Comment
func (ctx *TaskMigration) Up() {
	migrations.Create(orm.DefaultDatabaseName, "tasks", func(table *migrations.Table) {
		table.Increment("id")
		table.BigInteger("user_id")
		table.TimeStamp("created_at").Current()
		table.String("title")
		table.String("description")
		table.Boolean("complete")
	})
}

// Comment
func (ctx *TaskMigration) Down() {
	migrations.Drop("sqlite", "tasks")
}
