package connections

import (
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

var (
	HandlerCtx  *model.HandlerCtx
	EchoContext *echo.Echo
)
