package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		return c.File("public/home.htmx")
	})

	e.GET("/internship", func(c echo.Context) error {
		return c.File("public/internship.htmx")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
