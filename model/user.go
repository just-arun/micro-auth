package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	UserName  string    `json:"userName"`
	Password  string    `json:"password"`
	Roles     []Role    `json:"roles" gorm:"many2many:user_role;foreignKey:ID"`
	Apps      []App     `json:"apps" gorm:"many2many:user_app;foreignKey:ID"`
	Profile   Profile   `json:"profile" gorm:"foreignKey:ID;references:ID"`
	CreatedAt time.Time `json:"createdAt"`
}
