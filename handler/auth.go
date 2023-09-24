package handler

import (
	"encoding/json"
	"errors"
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
		fmt.Println(0)

		// get app settings
		general, err := service.General().Get(ctx.DB)
		if err != nil {
			return err
		}
		fmt.Println(general)
		if !general.CanLogin {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"error": "logging is disabled CODE(1)",
			})
		}
		fmt.Println(1)

		// extracting user input
		var body *requestDto.Login
		if err := json.
			NewDecoder(
				c.Request().Body,
			).
			Decode(&body); err != nil {
			return err
		}

		filter := &model.User{Email: body.Email}
		// getting user form data
		user, err := service.User().GetOne(ctx.DB, filter)
		if err != nil {
			return err
		}
		fmt.Println(3)
		if user.Type != model.UserTypeVerified {
			fmt.Println("USER_TYPE", user.Type)
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"error": "logging is disabled CODE(2)",
			})
		}
		fmt.Println(4)
		bodyHash, err := util.Password().Hash(body.Password)
		if err != nil {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"error": err.Error(),
			})
		}
		fmt.Println(bodyHash)

		fmt.Println(5)
		if !util.Password().Compare(user.Password, body.Password) {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"error": "invalid credentials",
			})
		}
		fmt.Println(6)

		roles := []string{}

		for _, v := range user.Roles {
			roles = append(roles, v.Name)
		}
		fmt.Println(7)

		resp, err := grpcClient.UserSession().SetUserSession(*ctx.GrpcClient, user.ID, roles)
		if err != nil {
			return err
		}
		fmt.Println(8)

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

		userID, err := service.User().CreateOne(ctx.DB, &user)
		if err != nil {
			return err
		}

		otp := util.NewOTP(6)

		_, err = grpcClient.OTPSession().SetOTP(*ctx.GrpcClient, otp, userID, model.OTPKeyRegisterVerify)
		if err != nil {
			return err
		}

		err = service.Mail().SendOtp(ctx.DB, user.Email, `your OTP: `+otp+``)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"ok":  1,
			"otp": otp,
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

func (a Auth) ResendRegisterVerifyOTP(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body requestDto.GetNewVerifyUserOTP
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

		otp := util.NewOTP(6)

		_, err = grpcClient.OTPSession().SetOTP(*ctx.GrpcClient, otp, user.ID, model.OTPKeyRegisterVerify)
		if err != nil {
			return err
		}

		err = service.Mail().SendOtp(ctx.DB, user.Email, `your OTP: `+otp+``)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "user verified",
		})
	}
}

func (a Auth) RegisterVerify(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body requestDto.VerifyUser
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

		exists, err := grpcClient.OTPSession().GetOTP(*ctx.GrpcClient, body.OTP, user.ID, model.OTPKeyRegisterVerify)
		if err != nil {
			return err
		}

		if !exists.Ok {
			return errors.New("invalid otp")
		}

		user.Type = model.UserTypeVerified

		err = service.User().UpdateVerify(ctx.DB, user.ID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "user verified",
		})
	}
}
