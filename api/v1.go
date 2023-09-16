package api

import (
	"fmt"

	"github.com/just-arun/micro-auth/boot"
	"github.com/just-arun/micro-auth/connections"
	"github.com/just-arun/micro-auth/handler"
	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/util"
	"github.com/labstack/echo/v4"
)

func apiV1(e *echo.Echo, g *echo.Group, environment, port, context string) {

	env := &model.Env{}
	n := fmt.Sprintf(".env.%v", environment)
	util.GetEnv(n, context, &env)
	ctx := &model.HandlerCtx{}
	ctx.Env = env
	pDb := boot.PostgresDB(env.DB.Uri)
	ctx.DB = pDb
	connections.DB = pDb
	userSession := boot.Redis(
		env.UserSession.Address,
		env.UserSession.Password,
		env.UserSession.DB,
		"UserSession",
	)
	ctx.UserSession = userSession
	generalSession := boot.Redis(
		env.GeneralSession.Address,
		env.GeneralSession.Password,
		env.GeneralSession.DB,
		"GeneralSession",
	)
	ctx.GeneralSession = generalSession

	v1 := g.Group("/v1")
	handler.Role(v1, ctx)

	serverPort := fmt.Sprintf(":%v", port)
	e.Logger.Fatal(
		e.Start(serverPort),
	)
}
