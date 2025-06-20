package controllers

import (
	"github.com/lucas11776-golang/http"
)

// Todo List View
func PageNotFound(req *http.Request, res *http.Response) *http.Response {
	return res.View("404", http.ViewData{})
}
