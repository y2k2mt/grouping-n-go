package handler

import (
	e "github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"github.com/y2k2mt/grouping-n-go/errors"
	"github.com/y2k2mt/grouping-n-go/infra"
	"github.com/y2k2mt/grouping-n-go/model"
	"net/http"
)

type GroupHandler struct {
	GroupRepository infra.GroupRepository
}

func (g *GroupHandler) GetGroup(c echo.Context) error {
	group, err := model.GetGroup("id", g.GroupRepository)
	if err != nil {
		if e.Is(err, errors.NoGroup) {
			return c.NoContent(http.StatusNotFound)
		} else {
			return c.NoContent(http.StatusInternalServerError)
		}
	}
	return c.JSON(http.StatusOK, group)
}
