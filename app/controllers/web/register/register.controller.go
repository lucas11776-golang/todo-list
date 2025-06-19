package register

import (
	"server/app/services/authentication"

	"github.com/lucas11776-golang/http"
)

// Register  View
func Index(req *http.Request, res *http.Response) *http.Response {
	return res.View("authentication.register", http.ViewData{})
}

// Register User
func Store(req *http.Request, res *http.Response) *http.Response {
	user, err := authentication.Register(req.Validator.Values())

	if err != nil {
		return res.Back().WithError("register_error", err.Error())
	}

	res.Session.Set("user_id", user.ID)

	return res.Redirect("todos")
}
