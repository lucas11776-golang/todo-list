package login

import (
	"server/app/services/authentication"

	"github.com/lucas11776-golang/http"
)

// Login Page
func Index(req *http.Request, res *http.Response) *http.Response {
	return res.View("authentication.login", http.ViewData{})
}

// Login User
func Store(req *http.Request, res *http.Response) *http.Response {
	user, err := authentication.Login(
		req.Validator.Value("email"),
		req.Validator.Value("password"),
	)

	if err != nil {
		return res.Back().WithError("login_error", err.Error())
	}

	res.Session.Set("user_id", user.ID)

	return res.Redirect("todos")
}
