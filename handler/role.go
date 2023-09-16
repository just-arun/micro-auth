package handler

import (
	"net/http"
	"strconv"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/service"
	"github.com/labstack/echo/v4"
)

type role struct{}

func Role(r *echo.Group, ctx *model.HandlerCtx) {
	st := &role{}
	r.GET("/role", st.GetNames(ctx))
}

func (r role) GetNames(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := service.Role().GetNames(ctx.DB)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"roles": data,
		})
	}
}

func (r role) GetOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return err
		}
		data, err := service.Role().GetOne(ctx.DB, uint(id))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"role": data,
		})
	}
}

func (r role) AddRole(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := new(model.Role)
		if err := c.Bind(role); err != nil {
			return err
		}
		return service.Role().Add(ctx.DB, *role)
	}
}
