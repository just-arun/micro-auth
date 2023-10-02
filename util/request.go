package util

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type request struct{ c echo.Context }

func Req(c echo.Context) *request {
	return &request{c}
}

func (c *request) Body(data interface{}) error {
	return json.NewDecoder(c.c.Request().Body).Decode(&data)
}
