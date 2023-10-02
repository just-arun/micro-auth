package responsedto

import "github.com/just-arun/micro-auth/model"

type User struct {
	ID       uint         `json:"id"`
	Email    string       `json:"email"`
	UserName string       `json:"userName"`
	Roles    []model.Role `json:"roles"`
}
