package routes

import (
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func Access(r *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.Access{}
	rout := r.Group("/access")
	rout.POST("/", st.AddOne(ctx))
	rout.GET("/", st.GetAll(ctx))
	rout.DELETE("/:id", st.DeleteOne(ctx))
}
