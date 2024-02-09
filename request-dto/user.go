package requestdto

import "github.com/just-arun/micro-auth/model"

type AddUser struct {
	Email    string         `json:"email"`
	UserName string         `json:"userName"`
	Type     model.UserType `json:"type"`
	Roles    []model.Role   `json:"roles"`
}

type UserList struct {
	Data []AddUser `json:"data"`
}
