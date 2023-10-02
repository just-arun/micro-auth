package handler

import (
	"net/http"
	"strconv"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/service"
	"github.com/just-arun/micro-auth/util"
	"github.com/labstack/echo/v4"
)

type Profile struct{}

func (a General) GetMyProfile(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		user, err := service.
			User().
			GetOne(ctx.DB, &model.User{ID: uint(id)})
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"profile": user,
		})
	}
}
