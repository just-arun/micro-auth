package routes

import (
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func Auth(rout *echo.Group, ctx *model.HandlerCtx) {
	st := &handler.Auth{}
	rout.POST("/login", st.Login(ctx))
	rout.POST("/register", st.Register(ctx))
	rout.GET("/public-key", st.GetPublicKey(ctx))
	rout.POST("/forgot-password", st.ForgotPassword(ctx))
	rout.POST("/update-password", st.UpdatePassword(ctx))
	rout.POST("/reset-password", st.ResetPassword(ctx))
}
