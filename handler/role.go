package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/service"
	"github.com/labstack/echo/v4"
)

type Role struct{}

func (r Role) GetNames(ctx *model.HandlerCtx) echo.HandlerFunc {
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

func (r Role) GetOne(ctx *model.HandlerCtx) echo.HandlerFunc {
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

func (r Role) AddRole(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		var role model.Role

		if err := json.
			NewDecoder(
				c.Request().Body,
			).
			Decode(&role); err != nil {
			return err
		}

		err := service.Role().Add(ctx.DB, role)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data": "Role Created",
		})
	}
}
