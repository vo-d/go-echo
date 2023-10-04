package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Cat struct {
	Name string
	Type string
}

func addCats(c echo.Context) error {
	// initialize objects

	cat := Cat{}

	// read the request body and put it to b
	// use io.ReadAll()
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("failed to read the request body: %v", err)
		return c.String(http.StatusInternalServerError, "")
	}
	// Unmarshal() decode the request body from the json to a slice, then store it to the address
	// If &cat is not the address, it will return an error to object. If &cat is the address, it will return nil to object
	// equivalent to JSON.parse() in ExpressJS
	object := json.Unmarshal(b, &cat)
	log.Print(cat)
	if object != nil {
		log.Printf("failed to parse the request body: %v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	// %#v parse the value in a default format
	log.Printf("this is your cat: %#v", cat)
	c.Request().Body.Close()

	return c.String(http.StatusOK, "We got your cat")

}

func main() {
	fmt.Println("Hello to the server")

	e := echo.New()

	e.POST("/cats", addCats)

	e.Start(":8000")
}
