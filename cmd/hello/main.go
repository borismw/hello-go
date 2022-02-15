package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const srvPort = ":8080"
const healthPort = ":1323"

func main() {

	// Separate Greetings and k8s health route onto different ports to avoid accidentally exposing health routes in k8s
	srv := SetupGreetingsServer()
	health := SetupHealthServer()

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
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		srv.Logger.Fatal(err)
	}
	if err := health.Shutdown(ctx); err != nil {
		health.Logger.Fatal(err)
	}
}
