package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/pubsub"
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

		allData, err := service.ServiceMap().GetMany(ctx.DB)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = pubsub.Publisher().ChangeServiceMap(ctx.NatsConnection, allData)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
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
			return util.Res(c).SendError(http.StatusConflict, fmt.Errorf("Not Found"))
		}
		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	}
}

func (st ServiceMap) GetMany(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := service.ServiceMap().GetMany(ctx.DB)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"serviceMap": data,
		})
	}
}

func (st ServiceMap) UpdateMany(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		// var body
		// data, err := service.ServiceMap().GetMany(ctx.DB)
		// if err != nil {
		// 	return util.Res(c).SendError(http.StatusConflict, err)
		// }
		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			// "serviceMap": data,
		})
	}
}

// func (st ServiceMap) DeleteOne(ctx *model.HandlerCtx) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		pId := c.Param("id")
// 		id, err := strconv.ParseUint(pId, 10, 32)
// 		if err != nil {
// 			return util.Res(c).SendError(http.StatusConflict, err)
// 		}
// 		err = service.ServiceMap().DeleteOne(ctx.DB, uint(id))
// 		if err != nil {
// 			return util.Res(c).SendError(http.StatusConflict, err)
// 		}
// 		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
// 			"ok": true,
// 		})
// 	}
// }
