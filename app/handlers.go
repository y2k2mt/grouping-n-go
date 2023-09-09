package app

import (
	"github.com/y2k2mt/grouping-n-go/handler"
	"github.com/y2k2mt/grouping-n-go/infra"
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

	groupRepository := infra.GroupRepository{DB: db}

	healthzHandler := handler.HealthzHandler{}
	groupHandler := handler.GroupHandler{GroupRepository: groupRepository, Log: logger}

	return Handlers{
		GroupHandler:   groupHandler,
		HealthzHandler: healthzHandler,
	}
}
