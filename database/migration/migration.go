package main

import (
	"fmt"
	"reflect"
	"server/database"
	"server/models"
	"time"

	"github.com/lucas11776-golang/http/utils/env"
	"github.com/lucas11776-golang/orm"
)

func main() {
	env.Load(".env")

	migration := database.Setup().Migration()

	for _, model := range tables() {
		_type := reflect.ValueOf(model)

		if err := migration.Migrate(orm.Models{model}); err != nil {
			fmt.Printf("Migration Failed (%s): %s\r\n", _type.Type().Name(), time.Now().Format(time.DateTime))

			continue
		}

		fmt.Printf("Migrated (%s): %s\r\n", _type.Type().Name(), time.Now().Format(time.DateTime))
	}
}

// Comment
func tables() orm.Models {
	return orm.Models{
		models.User{},
		models.Task{},
	}

}
