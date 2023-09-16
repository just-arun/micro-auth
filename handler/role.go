package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/service"
	"github.com/labstack/echo/v4"
)

type role struct{}

func Role(r *echo.Group, ctx *model.HandlerCtx) {
	st := &role{}
	rout := r.Group("/role")
	rout.GET("/", st.GetNames(ctx))
	rout.GET("/:id", st.GetOne(ctx))
	rout.POST("/", st.AddRole(ctx))
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
		var role model.Role

		if err := json.
			NewDecoder(
				c.Request().Body,
			).
			Decode(&role); err != nil {
			return err
		}

		return service.Role().Add(ctx.DB, role)
	}
}
