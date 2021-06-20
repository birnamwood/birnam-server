package main

import (
	"flag"
	"os"

	"github.com/birnamwood/birnam-server/init/config"
	"github.com/birnamwood/birnam-server/init/router"
	"github.com/birnamwood/birnam-server/init/router/database"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// 環境設定取得 flag.String(<パラメータ名>, <デフォルト値>, <パラメータの説明>)
	env := flag.String("e", "production", "動作環境名")
	//パラメータを渡してconfigの初期化を行う
	config.Init(*env, "environment")
	database.Init()
	router.Init(port)

}
