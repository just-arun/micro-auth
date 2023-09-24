package routes

import (
	"github.com/just-arun/micro-auth/acl"
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/middleware"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func Access(r *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.Access{}
	rout := r.Group("/access")
	rout.POST("/", st.AddOne(ctx), middleware.Auth(ctx, acl.ACLAccessAddOne))
	rout.GET("/", st.GetAll(ctx), middleware.Auth(ctx, acl.ACLAccessGetAll))
	rout.DELETE("/:id", st.DeleteOne(ctx), middleware.Auth(ctx, acl.ACLAccessDeleteOne))
}
