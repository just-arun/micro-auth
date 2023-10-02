package api

import "github.com/labstack/echo/v4"

func Run(context, env, port string, noServer ...bool) {
	e := echo.New()
	a := e.Group("/api")
	apiV1(e, a, env, port, context, noServer...)
}
