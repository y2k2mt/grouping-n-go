package app

import (
	"github.com/y2k2mt/grouping-n-go/infra"
	"github.com/y2k2mt/grouping-n-go/internal/handler"
)

type Handlers struct {
	HealthzHandler handler.HealthzHandler
	GroupHandler   handler.GroupHandler
}

func NewHandlers() Handlers {

	// Initialize all application instances.(Constructor injection)
	// DI container or Functional Option Pattern are too heavy...
	logger := GetAppLogger()
	db := infra.GetDatabase()

	groupingDatasource := infra.GroupingDatasource{DB: db}

	healthzHandler := handler.HealthzHandler{}
	groupHandler := handler.GroupHandler{Groups: groupingDatasource, Log: logger}

	return Handlers{
		GroupHandler:   groupHandler,
		HealthzHandler: healthzHandler,
	}
}
