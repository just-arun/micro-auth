package model

import "gorm.io/gorm"

type HandlerCtx struct {
	Env *Env
	DB  *gorm.DB
}
