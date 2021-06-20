package router

import (
	"net/http"

	"github.com/birnamwood/birnam-server/pkg/interface/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(port string) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = false
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://birnamwood.github.io", "http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))
	// ルーティング設定
	e.GET("/", handler.GetTest)
	e1 := e.Group("/api")
	{
		e1.POST("/test", handler.GetTest2)
	}
	e.Start(":" + port)

}
