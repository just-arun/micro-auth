package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/service"
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
			return err
		}

		err := service.Access().AddOne(ctx.DB, access)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data": "Access Created",
		})
	}
}

func (r Access) DeleteOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return err
		}

		err = service.Access().DeleteOne(ctx.DB, uint(id))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": "Access Deleted",
		})
	}
}

func (r Access) GetAll(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {

		data, err := service.Access().GetAll(ctx.DB)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"accesses": data,
		})
	}
}
