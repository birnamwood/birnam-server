package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetTest 疎通確認
func GetTest(c echo.Context) error {
	v := map[string]string{"Result": "Welcome!!"}
	return c.JSON(http.StatusOK, v)
}

// GetTest2 疎通確認
func GetTest2(c echo.Context) error {
	v := map[string]string{"Result": "OK2"}
	return c.JSON(http.StatusOK, v)
}
