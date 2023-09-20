package handler

import (
	"fmt"
	"net/http"

	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/util"
	"github.com/labstack/echo/v4"
)

type Auth struct{}

func (a Auth) Login(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (a Auth) Register(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
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
