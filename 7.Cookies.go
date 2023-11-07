package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "you're on the secret admin page")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you're on the secret admin page")

}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// http://localhost:8000/login?username=jack&password=12345
	// Check username and password against DB after hashing it
	if username == "jack" && password == "12345" {
		cookie := new(http.Cookie)

		//this is the same
		//cookie := &http.Cookie{}
		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		return c.String(http.StatusOK, "You were logged in!")
	}

	return c.String(http.StatusUnauthorized, "Your username or password were incorrect")
}

// middleware to check cookie validity
func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check cookie in req with the key  "sessionID"
		cookie, err := c.Cookie("sessionID")

		// if there is no coookie, return error string
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "You don't have any cookies")
			}
			log.Println(err)
			return err
		}
		// if there is cookie with value "some_string", proceed to next
		if cookie.Value == "some_string" {
			return next(c)
		}

		// if there is cookie with value different from "some_string", return string
		return c.String(http.StatusUnauthorized, "You don't have the right cookie")
	}
}

func main() {
	fmt.Println("Hello to the server")

	e := echo.New()

	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")

	// localhost:8000/admin/main
	adminGroup.GET("/main", mainAdmin, middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// everytime visit cookie group, run checkCookie
	cookieGroup.Use(checkCookie)
	cookieGroup.GET("/main", mainCookie)

	e.GET("/login", login)

	e.Start(":8000")
}
