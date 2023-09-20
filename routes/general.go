package routes

import (
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func General(route *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.General{}
	rout := route.Group("/general")
	rout.POST("/", st.Create(ctx))
	rout.GET("/", st.Get(ctx))
	rout.GET("/:id", st.Update(ctx))
}
