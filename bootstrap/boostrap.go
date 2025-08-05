package bootstrap

import (
	"fmt"
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
	if envPath != "" {
		LoadEnv(envPath)
	}

	server := http.Server(env.Env("HOST"), env.EnvInt("PORT"))

	ConfigureOptions(server)
	ConfigureDatabase()
	ConfigureRoutes(server)

	return server
}

// Comment
func Server(envPath string) {
	server := Boot(envPath)

	fmt.Printf("\r\nRunning Server %s:%d\r\n", env.Env("HOST"), env.EnvInt("PORT"))

	server.Listen()
}

// Comment
func Migration(envPath string) {
	database.Migrate(envPath)
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
