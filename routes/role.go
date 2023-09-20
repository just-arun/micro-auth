package routes

import (
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func Role(r *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.Role{}
	rout := r.Group("/role")
	rout.GET("/", st.GetNames(ctx))
	rout.GET("/:id", st.GetOne(ctx))
	rout.POST("/", st.AddRole(ctx))
}
