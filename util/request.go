package util

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type request struct{ c echo.Context }

func Req(c echo.Context) *request {
	return &request{c}
}

func (c *request) Body(data interface{}) error {
	return json.NewDecoder(c.c.Request().Body).Decode(&data)
}

type Pagination struct {
	Limit int64 `json:"limit"`
	Skip  int64 `json:"skip"`
	Page  int64 `json:"page"`
	Total int64 `json:"total"`
}

func (c *request) Pagination() (*Pagination, error) {

	pagination := &Pagination{Limit: 0, Skip: 0, Page: 0, Total: 0}
	var err error

	pageStr := c.c.QueryParam("page")
	if len(pageStr) > 0 {
		pagination.Page, err = strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			return &Pagination{}, err
		}
	}

	if pagination.Page <= 0 {
		pagination.Page = 1
	}

	limitStr := c.c.QueryParam("limit")
	if len(limitStr) > 0 {
		pagination.Limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			return &Pagination{}, err
		}
	}

	fmt.Println(pagination.Limit)

	if pagination.Limit <= 0 {
		pagination.Limit = 10
	}

	if pagination.Limit >= 1 {
		pagination.Skip = (pagination.Limit * pagination.Page) - pagination.Limit
	}

	return pagination, nil
}
