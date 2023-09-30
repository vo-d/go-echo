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
	e.GET("/cat", getCats)

	e.Start(":8000")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	// %s is parse string
	return c.String(http.StatusOK, fmt.Sprintf("your cat name is %s\nand his type is: %s\n", catName, catType))
}
