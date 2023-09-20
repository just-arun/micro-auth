package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	grpcClient "github.com/just-arun/micro-auth/grpcClient"
	"github.com/just-arun/micro-auth/model"
	requestDto "github.com/just-arun/micro-auth/request-dto"
	"github.com/just-arun/micro-auth/service"
	"github.com/just-arun/micro-auth/util"
	"github.com/labstack/echo/v4"
)

type Auth struct{}

func (a Auth) Login(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		general, err := service.General().Get(ctx.DB)
		if err != nil {
			return err
		}
		if general.CanLogin {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"error": "logging is disabled CODE(1)",
			})
		}
		var body requestDto.Login
		if err := json.
			NewDecoder(
				c.Request().Body,
			).
			Decode(&body); err != nil {
			return err
		}

		user, err := service.User().GetOne(ctx.DB, &model.User{Email: body.Email})
		if err != nil {
			return err
		}
		if user.Type != model.UserTypeVerified {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"error": "logging is disabled CODE(2)",
			})
		}

		if !util.Password().Compare(user.Password, body.Password) {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"error": "invalid credentials",
			})
		}

		roles := []string{}

		for _, v := range user.Roles {
			roles = append(roles, v.Name)
		}

		resp, err := grpcClient.UserSession().SetUserSession(*ctx.GrpcClient, user.ID, roles)
		if err != nil {
			return err
		}

		if general.HttpOnlyCookie {
			c.SetCookie(&http.Cookie{
				Name:   "x-session",
				Value:  resp.Token,
				Path:   "/",
				Secure: true,
				MaxAge: int(general.AccessTokenExpiryTime),
			})
			c.SetCookie(&http.Cookie{
				Name:   "x-refresh",
				Value:  resp.Refresh,
				Path:   "/",
				Secure: true,
				MaxAge: int(general.RefreshTokenExpiryTime),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"daa": map[string]interface{}{
				"accessToken":  resp.Token,
				"refreshToken": resp.Refresh,
			},
		})
	}
}

func (a Auth) Register(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		general, err := service.General().Get(ctx.DB)
		if err != nil {
			return err
		}
		if general.CanRegister {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"error": "registration is disabled",
			})
		}
		var user model.User
		if err := json.
			NewDecoder(
				c.Request().Body,
			).
			Decode(&user); err != nil {
			return err
		}

		_, err = service.User().CreateOne(ctx.DB, &user)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"ok": 1,
		})
	}
}

func (a Auth) GetPublicKey(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		privateKey, err := util.Rsa().PrivateKeyFrom64(ctx.Env.Rsa.PrivateKey)
		if err != nil {
			fmt.Println(err)
			return err
		}
		publicKey, err := util.Rsa().GeneratePublicKeyBase64(privateKey)
		if err != nil {
			fmt.Println(err)
			return err
		}

		_, err = grpcClient.UserSession().SetUserSession(*ctx.GrpcClient, 2, []string{"admin", "analyst"})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"key": publicKey,
		})
	}
}

func (a Auth) ForgotPassword(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (a Auth) UpdatePassword(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (a Auth) ResetPassword(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
