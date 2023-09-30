package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func getCats(c echo.Context) error {

	// get a specific parameter from the url
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	// Param() will return the datatype of url
	// ex: http://localhost:8000/cats/string?name=luna&type=yellow
	dataType := c.Param("data")

	// evaluate the data
	if dataType == "string" {
		// %s is parse string
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is %s\nand his type is: %s\n", catName, catType))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "you need to lets us know if you want json or string data",
	})

}

func main() {
	fmt.Println("Hello to the server")

	e := echo.New()

	e.GET("/", hello)
	e.GET("/cats/:data", getCats)

	e.Start(":8000")
}
