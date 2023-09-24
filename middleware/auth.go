package middleware

import (
	"net/http"

	"github.com/just-arun/micro-auth/acl"
	grpcclient "github.com/just-arun/micro-auth/grpcClient"
	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
)

func Auth(ctx *model.HandlerCtx, routeAccessReferenceKey acl.ACL) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			result, err := grpcclient.UserSession().HaveAccess(*ctx.GrpcClient, c.Request(), string(routeAccessReferenceKey))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "unauthorized",
				})
			}
			if !result.Access {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "unauthorized",
				})
			}
			return next(c)
		}
	}
}
