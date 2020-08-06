package main

import (
	"github.com/bullrox/rapidcaptcha-server/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/bullrox/rapidcaptcha-server/docs"
)

func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set up basic auth with username=foo and password=bar
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username, password string, c echo.Context) (bool, error) {
			if username == "foo" && password == "bar" {
				return true, nil
			}
			return false, nil
		},
	}))

	// Route => handler
	e.POST("/foobar", api.FooBarHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
