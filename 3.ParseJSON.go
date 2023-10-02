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
	Name string `json:"name"`
	Type string `json:"type"`
}

func addCats(c echo.Context) error {
	// initialize objects
	cat := Cat{}

	// read the request body and put it to b
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	// Unmarshal() decode the request body from the json to a slice
	// equivalent to JSON.parse() in ExpressJS
	object := json.Unmarshal(b, &cat)
	if object != nil {
		log.Printf("failed unmarshaling in addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	// %#v parse the value in a default format
	log.Printf("this is your cat: %#v", cat)
	c.Request().Body.Close()

	return c.String(http.StatusOK, "we got your cat")

}

func main() {
	fmt.Println("Hello to the server")

	e := echo.New()

	e.POST("/cats", addCats)

	e.Start(":8000")
}
