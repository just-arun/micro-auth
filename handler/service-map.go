package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/just-arun/micro-auth/model"
	requestdto "github.com/just-arun/micro-auth/request-dto"
	"github.com/just-arun/micro-auth/service"
	"github.com/just-arun/micro-auth/util"
	"github.com/labstack/echo/v4"
)

type ServiceMap struct{}

func (st ServiceMap) Add(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		var serviceMap *requestdto.CreateServiceMap
		if err := json.NewDecoder(c.Request().Body).Decode(&serviceMap); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		err := service.ServiceMap().Add(ctx.DB, &model.ServiceMap{
			Key:   serviceMap.Key,
			Value: serviceMap.Value,
			Auth:  serviceMap.Auth,
		})
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = service.ServiceMap().PublishSitemap(ctx.DB, ctx.NatsConnection)
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		return util.Res(c).SendSuccess(http.StatusCreated, map[string]interface{}{
			"ok": true,
		})
	}
}

func (st ServiceMap) GetOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		data, err := service.ServiceMap().GetOne(ctx.DB, uint(id))
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		if data == nil {
			return util.Res(c).SendError(http.StatusConflict, fmt.Errorf("not found"))
		}
		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	}
}

func (st ServiceMap) GetMany(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		search := c.QueryParam("search")

		pagination, err := util.Req(c).Pagination()
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		data, err := service.ServiceMap().GetMany(ctx.DB, search, pagination)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"serviceMap": data,
		}, map[string]interface{}{
			"metaData": pagination,
		})
	}
}

func (st ServiceMap) UpdateOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		var body model.ServiceMap
		if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
			return util.Res(c).SendError(http.StatusBadRequest, err)
		}

		serviceMap, err := service.ServiceMap().GetOne(ctx.DB, uint(id))
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		if serviceMap.Default {

			if err := service.ServiceMap().UpdateOne(ctx.DB, uint(id), &model.ServiceMap{
				Value:   body.Value,
				Auth:    body.Auth,
				Key:     serviceMap.Key,
				Default: serviceMap.Default,
			}); err != nil {
				return util.Res(c).SendError(http.StatusInternalServerError, err)
			}

			err = service.ServiceMap().PublishSitemap(ctx.DB, ctx.NatsConnection)
			if err != nil {
				return util.Res(c).SendError(http.StatusInternalServerError, err)
			}

			return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
				"ok": true,
			})

		}

		err = service.ServiceMap().UpdateOne(ctx.DB, uint(id), &body)

		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		err = service.ServiceMap().PublishSitemap(ctx.DB, ctx.NatsConnection)
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}
		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"ok": true,
		})
	}
}

func (st ServiceMap) DeleteOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		serviceMap, err := service.ServiceMap().GetOne(ctx.DB, uint(id))
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		if serviceMap.Default {
			return util.Res(c).SendError(http.StatusBadRequest, fmt.Errorf("can't delete default service map"))
		}

		err = service.ServiceMap().DeleteOne(ctx.DB, uint(id))
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = service.ServiceMap().PublishSitemap(ctx.DB, ctx.NatsConnection)
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"ok": true,
		})
	}
}
