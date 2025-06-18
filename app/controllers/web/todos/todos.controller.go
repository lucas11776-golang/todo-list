package todos

import (
	"github.com/lucas11776-golang/http"
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
	return res
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
