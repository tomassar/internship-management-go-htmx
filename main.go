package main

import (
	"net/http"
	"text/template"

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

	// Route for internships page
	e.GET("/internships", func(c echo.Context) error {
		internships := []map[string]string{
			{"title": "Internship 1", "description": "Description for Internship 1"},
			{"title": "Internship 2", "description": "Description for Internship 2"},
			{"title": "Internship 3", "description": "Description for Internship 3"},
		}

		tmpl, err := template.ParseFiles("public/internships.htmx")
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		err = tmpl.Execute(c.Response().Writer, internships)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
