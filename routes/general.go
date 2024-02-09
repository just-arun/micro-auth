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
	rout.GET("/history", st.GetMany(ctx))
	rout.PUT("/admin", st.UpdateAdmin(ctx)) // middleware.Auth(ctx, "auth.general-update-one"),
	rout.PUT("/dev", st.UpdateDev(ctx))     // middleware.Auth(ctx, "auth.general-update-one"),

}
