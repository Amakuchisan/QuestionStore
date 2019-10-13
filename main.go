package main

import (
	"github.com/Amakuchisan/QuestionBox/route"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := route.Init()
	e.Use(middleware.Logger())
	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":1323"))
}
