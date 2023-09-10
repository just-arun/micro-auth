package handler

import (
	"net/http"

	"github.com/just-arun/micro-auth/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type hello struct {
	env *model.Env
	db  *gorm.DB
}

func Hello(r *echo.Group) {
	h := &hello{}
	r.GET("/hello", h.HelloWorld())
}

func (h *hello) HelloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"data": "ok"})
	}
}
