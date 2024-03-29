package routes

import (
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func Access(r *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.Access{}
	rout := r.Group("/access")
	rout.POST("", st.AddOne(ctx)) // , middleware.Auth(ctx, acl.ACLAccessAddOne)

	rout.GET("", st.GetMany(ctx))                // middleware.Auth(ctx, acl.ACLAccessGetAll),
	rout.PUT("/:id/name", st.UpdateOneName(ctx)) // middleware.Auth(ctx, acl.ACLAccessGetAll),
	rout.GET("/:id/roles", st.LinkedRoles(ctx))

	// rout.DELETE("/:id", st.DeleteOne(ctx)) // middleware.Auth(ctx, acl.ACLAccessDeleteOne),

}
