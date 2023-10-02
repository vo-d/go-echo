package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello to the server")

	e := echo.New()

	e.GET("/", hello)

	e.Start(":8000")
}

// c echo.Context is (req, res) in ExpressJS but it can do both jobs
func hello(c echo.Context) error {

	// print Hello, world! to the page
	//return c.String(http.StatusOK, "Hello, world!")

	// fmt.Println("Hello, world!") = console.log("Hello, world!") in ExpressJS
	fmt.Println("Hello, world!")

	// log.Printf("Hello, world!") act similar to fmt.Println but it provides timestamps
	log.Printf("Hello, world!")

	// print nothing to the page
	return (c.String(http.StatusOK, ""))
}
