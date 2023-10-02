package requestdto

import "github.com/just-arun/micro-auth/model"

type AddAccessToRole struct {
	Accesses []model.Access `json:"accesses"`
}
