package util

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

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

func (c *response) SendSuccess(statusCode int, data map[string]interface{}, params ...map[string]interface{}) error {
	response := map[string]interface{}{
		"data": data,
	}

	fmt.Println("PARAMS", params, len(params))

	if len(params) > 0 {
		for k, v := range params[0] {
			response[k] = v
			fmt.Println("K: ", k)
			fmt.Println("V: ", v)
		}
	}

	return c.c.JSON(statusCode, response)
}
