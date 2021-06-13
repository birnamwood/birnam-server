package main

import (
	"os"

	"github.com/birnamwood/birnam-server/init/router"
	"github.com/birnamwood/birnam-server/init/router/database"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	database.Init()
	router.Init(port)

}
