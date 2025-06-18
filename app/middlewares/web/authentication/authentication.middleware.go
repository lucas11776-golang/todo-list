package authentication

import (
	"fmt"

	"github.com/lucas11776-golang/http"
)

// Check if user is guest - not logged int
func IsGuest(req *http.Request, res *http.Response, next http.Next) *http.Response {
	if req.Session.Get("user_id") != "" {
		return res.Redirect("/")
	}

	return next()
}

// Check if user is authenticated - is logged int
func IsAuth(req *http.Request, res *http.Response, next http.Next) *http.Response {
	if req.Session.Get("user_id") == "" {
		return res.Redirect(fmt.Sprintf("login?redirect=%s", req.Path()))
	}

	return next()
}
