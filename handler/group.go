package handler

import (
	e "github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"github.com/y2k2mt/grouping-n-go/errors"
	"github.com/y2k2mt/grouping-n-go/infra"
	"github.com/y2k2mt/grouping-n-go/model"
	"go.uber.org/zap"
	"net/http"
)

type GroupHandler struct {
	GroupRepository infra.GroupRepository
	Log             *zap.Logger
}

func (g *GroupHandler) GetGroup(c echo.Context) error {
	group, err := model.GetGroup(model.GroupId{Id: c.Param("id")}, g.GroupRepository)
	if err != nil {
		if e.Is(err, errors.NoGroup) {
			return c.NoContent(http.StatusNotFound)
		} else {
			if g.Log != nil {
				g.Log.Warn("failed to fetch group %v", zap.Error(err))
			}
			return c.NoContent(http.StatusInternalServerError)
		}
	}
	return c.JSON(http.StatusOK, group)
}
