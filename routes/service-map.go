package routes

import (
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func ServiceMap(r *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.ServiceMap{}
	rout := r.Group("/service-map")
	rout.GET("", st.GetMany(ctx))
	rout.GET("/:id", st.GetOne(ctx))
	rout.POST("", st.Add(ctx))
	rout.DELETE("/:id", st.DeleteOne(ctx))
}
