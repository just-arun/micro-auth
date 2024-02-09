package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	grpcclient "github.com/just-arun/micro-auth/grpcClient"
	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/service"
	"github.com/just-arun/micro-auth/util"
	"github.com/labstack/echo/v4"
)

type Role struct{}

func (r Role) GetNames(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := service.Role().GetNames(ctx.DB)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"roles": data,
		})
	}
}

func (r Role) GetOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		data, err := service.Role().GetOne(ctx.DB, uint(id))
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"role": data,
		})
	}
}

func (r Role) AddRole(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		var role model.Role

		if err := json.
			NewDecoder(
				c.Request().Body,
			).
			Decode(&role); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err := service.Role().Add(ctx.DB, &role)

		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		err = grpcclient.Role().SetRole(*ctx.GrpcClient, &role)
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		return util.Res(c).SendSuccess(http.StatusCreated, map[string]interface{}{
			"data": map[string]uint{
				"id": role.ID,
			},
		})
	}
}

func (r Role) UpdateAccesses(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		type Body struct {
			Accesses []model.Access `json:"accesses"`
		}

		var body Body

		if err := json.
			NewDecoder(
				c.Request().Body,
			).
			Decode(&body); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = service.Role().UpdateAccesses(ctx.DB, uint(id), body.Accesses)
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		err = r.updateRoleInSession(ctx, uint(id))
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"data": fmt.Sprintf("Access updated for role %v", id),
		})
	}
}

func (r Role) AddAccess(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		var access model.Access

		if err := json.
			NewDecoder(
				c.Request().Body,
			).
			Decode(&access); err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = service.Role().AddAccess(ctx.DB, uint(id), &access)

		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = r.updateRoleInSession(ctx, uint(id))
		if err != nil {
			return util.Res(c).SendError(http.StatusInternalServerError, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"data": "Access added to role",
		})
	}
}

func (r Role) DeleteOne(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		id, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = service.Role().DeleteOne(ctx.DB, uint(id))

		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"data": "role deleted",
		})
	}
}

func (r Role) RemoveOneAccess(ctx *model.HandlerCtx) echo.HandlerFunc {
	return func(c echo.Context) error {
		pId := c.Param("id")
		roleID, err := strconv.ParseUint(pId, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}
		pId0 := c.Param("accessID")
		accessID, err := strconv.ParseUint(pId0, 10, 32)
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		err = service.Role().RemoveOneAccess(ctx.DB, uint(roleID), uint(accessID))
		
		if err != nil {
			return util.Res(c).SendError(http.StatusConflict, err)
		}

		return util.Res(c).SendSuccess(http.StatusOK, map[string]interface{}{
			"message": "access removed",
		})
	}
}


func (r Role) updateRoleInSession(ctx *model.HandlerCtx, id uint) error {
	role, err := service.Role().GetOne(ctx.DB, id)

	if err != nil {
		return err
	}

	err = grpcclient.Role().SetRole(*ctx.GrpcClient, role)
	if err != nil {
		return err
	}
	fmt.Println("Role updated in session")
	return nil
}