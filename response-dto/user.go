package responsedto

import (
	"github.com/just-arun/micro-auth/model"
)

type User struct {
	ID       uint         `json:"id"`
	Email    string       `json:"email"`
	UserName string       `json:"userName"`
	Roles    []model.Role `json:"roles"`
}

type SlimUser struct {
	ID       uint              `json:"id" gorm:"primaryKey"`
	Email    string            `json:"email" gorm:"uniqueIndex"`
	UserName string            `json:"userName"`
	Type     model.UserType    `json:"type"`
	Roles    []GetAllRolesName `json:"roles"`
}
