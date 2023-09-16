package model

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type HandlerCtx struct {
	Env            *Env
	DB             *gorm.DB
	UserSession    *redis.Client
	GeneralSession *redis.Client
}
