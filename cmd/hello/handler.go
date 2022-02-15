package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strings"
)

func GetGreeting(c echo.Context) error {
	userid := c.Param("userid")

	if strings.TrimSpace(userid) == "" {
		return c.JSON(http.StatusBadRequest, "")
	}

	greeting := fmt.Sprintf("Hello %s!", userid)

	return c.JSON(http.StatusOK, greeting)
}

func SetupGreetingsServer() *echo.Echo {
	// Setup Greeting-API
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello/:userid", GetGreeting)

	return e
}

func SetupHealthServer() *echo.Echo {
	// Setup k8s health checks
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/readiness", func(c echo.Context) error { return c.JSON(http.StatusOK, "") })
	e.GET("/liveness", func(c echo.Context) error { return c.JSON(http.StatusOK, "") })

	return e
}
