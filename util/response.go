package util

import "github.com/labstack/echo/v4"

type response struct{ c echo.Context }

func Res(c echo.Context) *response {
	return &response{c}
}

func (c *response) SendError(statusCode int, err error) error {
	return c.c.JSON(statusCode, map[string]interface{}{
		"error": map[string]string{
			"message": err.Error(),
		},
	})
}

func (c *response) SendSuccess(statusCode int, data map[string]interface{}) error {
	return c.c.JSON(statusCode, map[string]interface{}{
		"data": data,
	})
}
