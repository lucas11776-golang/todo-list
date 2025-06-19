package tasks

import (
	"server/models"

	"github.com/lucas11776-golang/orm"
)

// Create user task
func CreateTask(userID int64, form map[string]string) (*models.Task, error) {
	return orm.Model(models.Task{}).
		Insert(orm.Values{
			"user_id":     userID,
			"title":       form["title"],
			"description": form["description"],
		})
}
