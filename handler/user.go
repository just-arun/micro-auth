package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/just-arun/micro-auth/model"
	requestdto "github.com/just-arun/micro-auth/request-dto"
	"github.com/just-arun/micro-auth/service"
	"github.com/just-arun/micro-auth/util"
	"github.com/labstack/echo/v4"
)

type User struct{}

func (h User) AddMultipleUser(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {

		var body requestdto.UserList
		if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		users := []*model.User{}

		for _, v := range body.Data {
			users = append(users, &model.User{
				Email:    v.Email,
				UserName: v.UserName,
				Password: util.RandStringRunes(12),
				Type:     model.UserTypeUnVerify,
			})
		}

		err := service.User().CreateMultiple(ctx.DB, users)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"ok": true,
		})
	}
}

func (h User) GetOne(ctx *model.HandlerCtx) echo.HandlerFunc {
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
			"user": user,
		})
	}
}

func (h User) GetMany(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {

		user, err := service.
			User().
			GetMany(ctx.DB, &model.User{})
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"users": user,
		})
	}
}
