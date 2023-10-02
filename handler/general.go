package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/service"
	"github.com/just-arun/micro-auth/util"
	"github.com/labstack/echo/v4"
)

type General struct{}

func (a General) Create(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		var general *model.General
		if err := json.NewDecoder(c.Request().Body).Decode(&general); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		err := service.General().Create(ctx.DB, general)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).SendSuccess(http.StatusCreated, map[string]interface{}{
			"ok": true,
		})
	}
}

func (a General) Get(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := service.General().Get(ctx.DB)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"general": data,
		})
	}
}

func (a General) Update(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		var general *model.General
		if err := json.NewDecoder(c.Request().Body).Decode(&general); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		err = service.General().Update(ctx.DB, uint(id), general)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).SendSuccess(http.StatusCreated, map[string]interface{}{
			"ok": true,
		})
	}
}
