package home

import (
	"github.com/lucas11776-golang/http"
)

// Home Page View
func Index(req *http.Request, res *http.Response) *http.Response {
	return res.View("index", http.ViewData{})
}
