package routes

import (
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func User(route *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.User{}
	rout := route.Group("/users")

	rout.POST("", st.AddMultipleUser(ctx))    // middleware.Auth(ctx, "auth.general-get"),
	rout.GET("/:id", st.GetOne(ctx))          // middleware.Auth(ctx, "auth.general-get"),
	rout.GET("", st.GetMany(ctx))          // middleware.Auth(ctx, "auth.general-get"),

}
