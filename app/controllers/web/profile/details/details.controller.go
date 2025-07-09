package details

import "github.com/lucas11776-golang/http"

// Change Account Details Page
func Edit(req *http.Request, res *http.Response) *http.Response {
	return res.View("authentication.login", http.ViewData{})
}

// Update Account Details Page
func Update(req *http.Request, res *http.Response) *http.Response {
	return res.Back()
}
