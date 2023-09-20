package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/service"
	"github.com/labstack/echo/v4"
)

type ServiceMap struct{}

func (st ServiceMap) Add(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		var serviceMap *model.ServiceMap
		if err := json.NewDecoder(c.Request().Body).Decode(&serviceMap); err != nil {
			return err
		}
		err := service.ServiceMap().Add(ctx.DB, serviceMap)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"ok": true,
		})
	}
}

func (st ServiceMap) GetOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return err
		}
		data, err := service.ServiceMap().GetOne(ctx.DB, uint(id))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	}
}

func (st ServiceMap) GetMany(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := service.ServiceMap().GetMany(ctx.DB)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	}
}

func (st ServiceMap) UpdateOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return err
		}
		var serviceMap *model.ServiceMap
		if err := json.NewDecoder(c.Request().Body).Decode(&serviceMap); err != nil {
			return err
		}
		err = service.ServiceMap().UpdateOne(ctx.DB, uint(id), serviceMap)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"ok": true,
		})
	}
}

func (st ServiceMap) DeleteOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return err
		}
		err = service.ServiceMap().DeleteOne(ctx.DB, uint(id))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"ok": true,
		})
	}
}
