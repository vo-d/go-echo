package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "you're on the secret admin page")
}

func main() {
	fmt.Println("Hello to the server")

	e := echo.New()

	// group here is like the router in ExpressJs
	// can add unlimited middleware to anything
	g := e.Group("/admin", middleware.Logger())

	//or g.Use(middleware.Logger())
	// this middleware logs the server interaction

	// localhost:8000/admin/main
	g.GET("/main", mainAdmin, middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// g.GET("/main", mainAdmin, middleware.Logger())

	e.Start(":8000")
}
