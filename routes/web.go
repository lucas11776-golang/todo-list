package routes

import (
	"server/app/controllers/web/home"
	"server/app/controllers/web/login"
	"server/app/controllers/web/logout"
	"server/app/controllers/web/register"
	"server/app/controllers/web/todos"
	"server/app/controllers/web/todos/complete"
	"server/app/middlewares/web/authentication"
	"server/app/middlewares/web/tasks"
	"server/app/requests"

	"github.com/lucas11776-golang/http"
)

// Web routes
func web(route *http.Router) {
	route.Get("/", home.Index)
	route.Group("register", func(route *http.Router) {
		route.Get("/", register.Index)
		route.Post("/", register.Store, requests.RegisterFormRequest())
	}, authentication.IsGuest)
	route.Group("login", func(route *http.Router) {
		route.Get("/", login.Index)
		route.Post("/", login.Store, requests.LoginFormRequest())
	}, authentication.IsGuest)
	route.Group("todos", func(route *http.Router) {
		route.Get("/", todos.Index)
		route.Get("create", todos.Create)
		route.Post("/", todos.Store, requests.CreateTask())
		route.Group("{task}", func(route *http.Router) {
			route.Get("/", todos.View)
			route.Get("/edit", todos.Edit)
			route.Patch("/", todos.Update, requests.UpdateTask())
			route.Delete("/", todos.Delete)
			route.Group("complete", func(route *http.Router) {
				route.Post("/", complete.Store)
				route.Delete("/", complete.Destroy)
			})
		}, tasks.UserTask)
	}, authentication.IsAuth)
	route.Delete("logout", logout.Destroy)
}
