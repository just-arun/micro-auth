package routes

import (
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func General(route *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.General{}
	rout := route.Group("/general")

	// rout.POST("/", st.Create(ctx),
	// 	middleware.Auth(ctx, "auth.general-create"))

	rout.GET("", st.Get(ctx)) // middleware.Auth(ctx, "auth.general-get"),

	rout.PUT("/:id", st.Update(ctx)) // middleware.Auth(ctx, "auth.general-update-one"),

}
