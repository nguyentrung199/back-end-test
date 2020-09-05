package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Nguyeexn quoc trung!")
}
func Handlersignin(c echo.Context) error {
	return c.String(http.StatusOK, "Nguyeexn quoc trung!")
}
func Handlersignup(c echo.Context) error {
	return c.String(http.StatusOK, "Nguyeexn quoc trung!")
}