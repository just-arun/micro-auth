package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/just-arun/micro-auth/model"
	requestdto "github.com/just-arun/micro-auth/request-dto"
	responsedto "github.com/just-arun/micro-auth/response-dto"
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

		generalResponse := &responsedto.GeneralGetOne{}
		err = util.UpateOneStructWithOtherWithSameKeys(data, generalResponse)
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"general": generalResponse,
		})
	}
}

func (a General) GetMany(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(00000)
		pagination, err := util.Req(c).Pagination()
		if err != nil {
			return util.Res(c).SendError(http.StatusBadRequest, err)
		}

		data, err := service.General().GetMany(ctx.DB, pagination)
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"generals": data,
		}, map[string]interface{}{
			"metaData": pagination,
		})
	}
}

func (a General) UpdateAdmin(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload *requestdto.UpdateGeneralAdminPayload
		if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		data, err := service.General().Get(ctx.DB)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = util.UpateOneStructWithOtherWithSameKeys(payload, data)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		data.UpdatedBy = append(data.UpdatedBy, model.User{ID: 1})

		err = service.General().Update(ctx.DB, data)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).SendSuccess(http.StatusCreated, map[string]interface{}{
			"ok": true,
		})
	}
}

func (a General) UpdateDev(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload *requestdto.UpdateGeneralServicePayload
		if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		data, err := service.General().Get(ctx.DB)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = util.UpateOneStructWithOtherWithSameKeys(payload, data)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		data.UpdatedBy = append(data.UpdatedBy, model.User{ID: 1})

		err = service.General().Update(ctx.DB, data)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).SendSuccess(http.StatusCreated, map[string]interface{}{
			"ok": true,
		})
	}
}
