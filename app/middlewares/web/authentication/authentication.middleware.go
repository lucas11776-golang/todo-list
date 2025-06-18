package authentication

import (
	"github.com/lucas11776-golang/http"
)

// Check if user is guest - not logged int
func IsGuest(req *http.Request, res *http.Response, next http.Next) *http.Response {
	return next()
}

// Check if user is authenticated - is logged int
func IsAuth(req *http.Request, res *http.Response, next http.Next) *http.Response {
	return next()
}
