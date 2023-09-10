package main

import (
	"context"
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/y2k2mt/grouping-n-go/app"
	"github.com/y2k2mt/grouping-n-go/infra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	err := app.InitAppConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	config := app.GetAppConfig()
	log.Printf("Start the application in %v mode", config.Env)

	err = app.InitAppLogger(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	appLogger := app.GetAppLogger()
	defer appLogger.Sync()

	err = infra.InitDatabase(config.DatabaseUrl, appLogger)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(echozap.ZapLogger(appLogger))
	e.Use(middleware.Recover())

	handlers := app.NewHandlers()

	e.GET("/healthz", handlers.HealthzHandler.Healthz)
	e.POST("/grouping", handlers.GroupHandler.CreateGroup)
	e.GET("/grouping/:id", handlers.GroupHandler.GetGroup)

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
