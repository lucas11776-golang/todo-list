package main

import (
	"fmt"
	"server/bootstrap"

	"github.com/lucas11776-golang/http/utils/env"
)

func main() {
	server := bootstrap.Boot(".env")

	fmt.Printf("\r\nRunning Server %s:%d\r\n", env.Env("HOST"), env.EnvInt("PORT"))

	server.Listen()
}
