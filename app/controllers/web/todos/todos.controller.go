package todos

import (
	"fmt"
	"server/app/services/tasks"

	"github.com/lucas11776-golang/http"
	"github.com/spf13/cast"
)

// Todo List View
func Index(req *http.Request, res *http.Response) *http.Response {
	tasks, err := tasks.Tasks(cast.ToInt64(req.Session.Get("user_id")))

	return res.View("todos.index", http.ViewData{
		"tasks_error": &err,
		"tasks":       &tasks,
	})
}

// Create Todo Page
func Create(req *http.Request, res *http.Response) *http.Response {
	return res.View("todos.create", http.ViewData{})
}

// Store Todo Item
func Store(req *http.Request, res *http.Response) *http.Response {
	task, err := tasks.CreateTask(
		cast.ToInt64(req.Session.Get("user_id")),
		req.Validator.Values(),
	)

	if err != nil {
		return res.Back().WithError("task_create_error", err.Error())
	}

	return res.Redirect(fmt.Sprintf("todos/%d", task.ID))
}

// View Todo Page
func View(req *http.Request, res *http.Response) *http.Response {
	task, err := tasks.Task(cast.ToInt64(req.Parameters.Get("task")))

	return res.View("todos.view", http.ViewData{
		"task_error": &err,
		"task":       task,
	})
}

// Edit Todo Page
func Edit(req *http.Request, res *http.Response) *http.Response {
	task, err := tasks.Task(cast.ToInt64(req.Parameters.Get("task")))

	return res.View("todos.edit", http.ViewData{
		"task_error": &err,
		"task":       task,
	})
}

// Update Todo
func Update(req *http.Request, res *http.Response) *http.Response {
	err := tasks.UpdateTask(
		cast.ToInt64(req.Parameters.Get("task")),
		req.Validator.Values(),
	)

	if err != nil {
		return res.Back().WithError("task_error", err.Error())
	}

	return res.Back()
}

// Delete Todo
func Delete(req *http.Request, res *http.Response) *http.Response {
	err := tasks.DeleteTask(cast.ToInt64(req.Parameters.Get("task")))

	if err != nil {
		return res.Back().WithError("task_error", err.Error())
	}

	return res.Redirect("todos")
}
