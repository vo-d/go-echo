package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello to the server")

	e := echo.New()

	e.GET("/", hello)

	e.GET("/", hello)

	e.Start(":8000")
}

// c echo.Context
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}
