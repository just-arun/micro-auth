package routes

import (
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func Auth(rout *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.Auth{}
	r := rout.Group("/auth")
	r.POST("/login", st.Login(ctx))
	r.POST("/register", st.Register(ctx))
	r.GET("/public-key", st.GetPublicKey(ctx))
	r.POST("/forgot-password", st.ForgotPassword(ctx))
	r.POST("/update-password", st.UpdatePassword(ctx))
	r.POST("/reset-password", st.ResetPassword(ctx))
}
