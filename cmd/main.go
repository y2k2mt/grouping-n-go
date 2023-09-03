package main

import (
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/y2k2mt/grouping-n-go/app"
	"github.com/y2k2mt/grouping-n-go/handler"
	"github.com/y2k2mt/grouping-n-go/infra"
	"log"
)

func main() {

	log.Print("Starging up application...")
	config, err := app.InitConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Initialized configuration: %v", config)

	err = app.InitLogger(config)
	if err != nil {
		log.Fatal(err.Error())
  }
  appLogger := app.GetZapLogger()
	defer appLogger.Sync()

	err = infra.InitDatabase(config.DatabaseUrl, appLogger)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(echozap.ZapLogger(appLogger))
	e.Use(middleware.Recover())

	healthz := handler.HealthzHandler{}
	group := handler.GroupHandler{GroupRepository: infra.GroupRepository{}}

	e.GET("/healthz", healthz.Healthz)
	e.GET("/group", group.GetGroup)

	e.Logger.Fatal(e.Start(":1323"))
}
