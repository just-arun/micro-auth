package handler

import (
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

type access struct{}

func Access(r *echo.Group, ctx *model.HandlerCtx) {
	st := access{}
	rout := r.Group("/access")
	rout.GET("/", st.GetNames(ctx))
}

func (r access) GetNames(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
