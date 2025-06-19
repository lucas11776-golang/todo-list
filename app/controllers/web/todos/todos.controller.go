package todos

import (
	"fmt"
	"server/app/services/tasks"

	"github.com/lucas11776-golang/http"
	"github.com/spf13/cast"
)

// Todo List View
func Index(req *http.Request, res *http.Response) *http.Response {
	return res.View("todos.index", http.ViewData{})
}

// Create Todo Page
func Create(req *http.Request, res *http.Response) *http.Response {
	return res.View("todos.create", http.ViewData{})
}

// Store Todo Item
func Store(req *http.Request, res *http.Response) *http.Response {
	task, err := tasks.CreateTask(
		cast.ToInt64(req.Session.Get("user-id")),
		req.Validator.Values(),
	)

	if err != nil {
		return res.Back().WithError("create_task_error", err.Error())
	}

	return res.Redirect(fmt.Sprintf("todos/%d", task.ID))
}

// View Todo Page
func View(req *http.Request, res *http.Response) *http.Response {
	return res.View("todos.view", http.ViewData{})
}

// Edit Todo Page
func Edit(req *http.Request, res *http.Response) *http.Response {
	return res.View("todos.edit", http.ViewData{})
}

// Update Todo
func Update(req *http.Request, res *http.Response) *http.Response {
	return res
}

// Delete Todo
func Delete(req *http.Request, res *http.Response) *http.Response {
	return res
}
