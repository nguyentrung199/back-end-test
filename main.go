package main

import (
	"Back_end/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handler.Home)
	e.GET("/user", handler.Handlersignin)
	e.GET("/info", handler.Handlersignup)
	e.Logger.Fatal(e.Start(":3000"))
}

