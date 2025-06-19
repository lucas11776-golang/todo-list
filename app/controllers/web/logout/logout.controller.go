package logout

import (
	"github.com/lucas11776-golang/http"
)

// Delete User Session
func Destroy(req *http.Request, res *http.Response) *http.Response {
	res.Session.Clear()

	return res.Redirect("/")
}
