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

func main() {
	fmt.Println("Hello to the server")

	e := echo.New()

	// group here is like the router in ExpressJs
	// this middleware logs the server  interaction
	g := e.Group("/admin", middleware.Logger())
	//or g.Use(middleware.Logger())

	// localhost:8000/admin/main
	g.GET("/main", mainAdmin)

	e.Start(":8000")

}
