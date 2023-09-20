package handler

import (
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

type Hello struct{}

func (h *Hello) SayHello(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"ok": true,
		})
	}
}
