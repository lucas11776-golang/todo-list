package main

import (
	"os"
	"server/bootstrap"
	"strings"
)

const (
	Env = ".env"
)

func main() {
	if len(os.Args) >= 2 {
		switch strings.ToUpper(os.Args[1]) {
		case "SERVER":
			bootstrap.Server(Env)

		case "MIGRATION":
			bootstrap.Migration(Env)
		}
	}
}
