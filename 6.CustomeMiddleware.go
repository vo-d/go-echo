package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func mainAdmin(e echo.Context) error {
	return e.String(http.StatusOK, "you're on the secret admin page")
}

// in expressJs, the middleware use (req, res) for the parameters, execute functionality and return the control back to the parent method
// in echo, the middleware use a handler function for the paramenter and return a function that use (req, res) for the paramenter, execute functionality and return the control back to the parent method
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		return next(c)
	}
}

func main() {
	fmt.Println("Hello to the server")

	e := echo.New()

	e.Use(ServerHeader)

	// group here is like the router in ExpressJs
	// can add unlimited middleware to anything
	g := e.Group("/admin", middleware.Logger())

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// check in the db
		if username == "jack" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	//or g.Use(middleware.Logger())
	// this middleware logs the server interaction

	// localhost:8000/admin/main
	g.GET("/main", mainAdmin, middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// g.GET("/main", mainAdmin, middleware.Logger())

	e.Start(":8000")
}
