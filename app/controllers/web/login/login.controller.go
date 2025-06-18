package login

import (
	"github.com/lucas11776-golang/http"
)

// Login  View
func Index(req *http.Request, res *http.Response) *http.Response {
	return res.View("authentication.login", http.ViewData{})
}

// Login User
func Store(req *http.Request, res *http.Response) *http.Response {
	return res
}
