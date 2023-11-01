package routes

import (
	"net/http"

	"github.com/just-arun/micro-auth/acl"
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/middleware"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func User(route *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.User{}
	rout := route.Group("/users")

	rout.POST("", st.AddMultipleUser(ctx)) // middleware.Auth(ctx, "auth.general-get"),
	rout.GET("/:id", st.GetOne(ctx))       // middleware.Auth(ctx, "auth.general-get"),
	rout.GET("", st.GetMany(ctx))          // middleware.Auth(ctx, "auth.general-get"),
	rout.PUT("/:id", st.UpdateUserRole(ctx))
	rout.GET("/some/other/stuff", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "some stuff",
		})
	},
		middleware.Auth(ctx, acl.ACLTestEndpoint),
	)
}
