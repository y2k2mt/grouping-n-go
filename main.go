package main

import (
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/y2k2mt/grouping-n-go/app"
	"github.com/y2k2mt/grouping-n-go/infra"
	"log"
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
	e.GET("/group/:id", handlers.GroupHandler.GetGroup)

	e.Logger.Fatal(e.Start(":1323"))
}
