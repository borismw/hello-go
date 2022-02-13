package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

const srvPort = ":8080"
const healthPort = ":1323"

func getGreeting(c echo.Context) error {
	userid := c.Param("userid")
    greeting := fmt.Sprintf("Hello %s!", userid)

	return c.JSON(http.StatusOK, greeting)
}

func setupGreetingsServer() *echo.Echo {
	// Setup Greeting-API
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello/:userid", getGreeting)

	return e
}

func setupHealthServer() *echo.Echo {
	// Setup k8s health checks
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/readiness", func(c echo.Context) error { return c.JSON(http.StatusOK, "") })
	e.GET("/liveness", func(c echo.Context) error { return c.JSON(http.StatusOK, "") })

	return e
}

func main() {

	// Separate Greetings and k8s health route onto different ports to avoid accidentally exposing health routes in k8s
	srv := setupGreetingsServer()
	health := setupHealthServer()

	go func() {
	  	if err := srv.Start(srvPort); err != nil && err != http.ErrServerClosed {
			  srv.Logger.Fatal("Shutting down the server")
		}
	}()

	go func() {
		if err := health.Start(healthPort); err != nil && err != http.ErrServerClosed {
			health.Logger.Fatal("Shutting down the server")
		}
	}()

	// Allow for graceful shutdown. Recipe from here: https://echo.labstack.com/cookbook/graceful-shutdown/
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<- quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		srv.Logger.Fatal(err)
	}
	if err := health.Shutdown(ctx); err != nil {
		health.Logger.Fatal(err)
	}
}