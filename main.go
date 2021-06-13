package main

import (
	"os"

	"github.com/birnamwood/birnam-server/init/router"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router.Init(port)

}
