package bootstrap

import (
	"server/database"
	"server/routes"

	"github.com/lucas11776-golang/http"
	"github.com/lucas11776-golang/http/utils/env"
)

func LoadEnv(envPath string) {
	env.Load(envPath)
}

// Comment
func Boot(envPath string) *http.HTTP {
	LoadEnv(envPath)

	server := http.Server(env.Env("HOST"), env.EnvInt("PORT"))

	ConfigureOptions(server)
	ConfigureDatabase()
	ConfigureRoutes(server)

	return server
}

// Comment
func ConfigureOptions(server *http.HTTP) {
	server.SetStatic(env.Env("ASSETS"))                                    // Static Assets
	server.SetView(env.Env("VIEWS"), env.Env("VIEWS_EXTENSION"))           // Views
	server.Session([]byte(env.Env("SESSION_KEY"))).Secure(false).Path("/") // Session
}

func ConfigureDatabase() {
	database.Setup()
}

// Comment
func ConfigureRoutes(server *http.HTTP) {
	server.Route().Callback(routes.Routes)
}
