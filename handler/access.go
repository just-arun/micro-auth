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

type Access struct{}

func (r Access) AddOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		var access model.Access

		if err := json.
			NewDecoder(
				c.Request().Body,
			).
			Decode(&access); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err := service.Access().AddOne(ctx.DB, access)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).SendSuccess(http.StatusCreated,
			map[string]interface{}{
				"data": "Access Created",
			})
	}
}

func (r Access) UpdateOneName(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		type NameStruct struct {
			Name string `json:"name"`
		}
		var access NameStruct

		if err := json.
			NewDecoder(
				c.Request().Body,
			).
			Decode(&access); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = service.Access().UpdateOneName(ctx.DB, uint(id), access.Name)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).
			SendSuccess(http.StatusCreated,
				map[string]interface{}{
					"data": "Access Updated",
				})
	}
}

// func (r Access) DeleteOne(ctx *model.HandlerCtx) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		pId := c.Param("id")
// 		id, err := strconv.ParseUint(pId, 10, 32)
// 		if err != nil {
// 			return util.Res(c).SendError(http.StatusConflict, err)
// 		}

// 		err = service.Access().DeleteOne(ctx.DB, uint(id))
// 		if err != nil {
// 			return util.Res(c).SendError(http.StatusConflict, err)
// 		}

// 		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
// 			"data": "Access Deleted",
// 		})
// 	}
// }

func (r Access) GetAll(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {

		data, err := service.Access().GetAll(ctx.DB)

		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"accesses": data,
		})
	}
}

func (r Access) GetMany(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {

		search := c.QueryParam("search")

		pagination, err := util.Req(c).Pagination()
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		data, err := service.Access().GetMany(ctx.DB, search, pagination)

		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		if data == nil {
			data = []model.Access{}
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"accesses": data,
		}, map[string]interface{}{
			"metaData": pagination,
		})
	}
}

func (r Access) LinkedRoles(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		resp, err := service.Access().LinkedRoles(ctx.DB, uint(id))
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}
		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"roles": resp,
		})
	}
}
