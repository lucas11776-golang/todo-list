package routes

import (
	"server/app/controllers"

	"github.com/lucas11776-golang/http"
)

// Routes
func Routes(route *http.Router) {
	route.Callback(web)
	route.Group("api", api)
	route.Fallback(controllers.PageNotFound)
}
