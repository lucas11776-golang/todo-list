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
			"complete":    false,
		})
}

// Create user task
func Tasks(userID int64) ([]*models.Task, error) {
	return orm.Model(models.Task{}).
		Where("user_id", "=", userID).
		Get()
}

// Create user task
func Task(taskID int64) (*models.Task, error) {
	return orm.Model(models.Task{}).
		Where("id", "=", taskID).
		First()
}

// Mark task as complete
func TaskCompleted(taskID int64) error {
	return orm.Model(models.Task{}).
		Where("id", "=", taskID).
		Update(orm.Values{"complete": true})
}

// Mark task as not complete
func TaskUncompleted(taskID int64) error {
	return orm.Model(models.Task{}).
		Where("id", "=", taskID).
		Update(orm.Values{"complete": false})
}

// Update Task
func UpdateTask(taskID int64, form map[string]string) error {
	return orm.Model(models.Task{}).
		Where("id", "=", taskID).
		Update(orm.Values{
			"title":       form["title"],
			"description": form["description"],
		})
}

// Update Task
func DeleteTask(taskID int64) error {
	return orm.Model(models.Task{}).
		Where("id", "=", taskID).
		Delete()
}
