package register

import (
	"github.com/lucas11776-golang/http"
)

// Register  View
func Index(req *http.Request, res *http.Response) *http.Response {
	return res.View("authentication.register", http.ViewData{})
}

// Register User
func Store(req *http.Request, res *http.Response) *http.Response {
	return res
}
