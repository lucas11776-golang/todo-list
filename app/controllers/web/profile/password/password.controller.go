package password

import "github.com/lucas11776-golang/http"

// Change Password Page
func Edit(req *http.Request, res *http.Response) *http.Response {
	return res.View("authentication.login", http.ViewData{})
}

// Update Password Page
func Update(req *http.Request, res *http.Response) *http.Response {
	return res.Back()
}
