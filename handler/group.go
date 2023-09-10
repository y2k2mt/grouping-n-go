package handler

import (
	"fmt"
	e "github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"github.com/y2k2mt/grouping-n-go/errors"
	"github.com/y2k2mt/grouping-n-go/infra"
	"github.com/y2k2mt/grouping-n-go/model"
	"go.uber.org/zap"
	"net/http"
)

type GroupHandler struct {
	Groups infra.GroupingDatasource
	Log    *zap.Logger
}

func (g *GroupHandler) CreateGroup(c echo.Context) error {
	candidates := new(model.Candidates)
	if err := c.Bind(candidates); err != nil {
		g.Log.Warn("%v", zap.Error(err))
		g.Log.Warn(fmt.Sprintf("%v", candidates))
		return c.String(http.StatusBadRequest, fmt.Sprintf("bad request %v", err))
	}
	identified, err := model.Grouping(*candidates)
	if err != nil {
		g.Log.Info("%v", zap.Error(err))
		if e.Is(err, errors.InsufficientGroupingMember) {
			return c.NoContent(http.StatusBadRequest)
		} else if e.Is(err, errors.InsufficientGroupingNumber) {
			return c.NoContent(http.StatusBadRequest)
		} else {
			if g.Log != nil {
				g.Log.Warn("failed to fetch group %v", zap.Error(err))
			}
			return c.NoContent(http.StatusInternalServerError)
		}
	}
	return c.JSON(http.StatusOK, identified)
}

func (g *GroupHandler) GetGroup(c echo.Context) error {
	group, err := model.GetGroup(model.GroupId{Id: c.Param("id")}, g.Groups)
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
