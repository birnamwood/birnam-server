package router

import (
	"github.com/birnamwood/birnam-server/pkg/interface/handler"
	"github.com/labstack/echo"
)

func Init(port string) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = false
	// ルーティング設定
	e.GET("/test", handler.GetTest)
	e1 := e.Group("/api")
	{
		e1.POST("/test", handler.GetTest2)
	}
	e.Start(":" + port)

}
