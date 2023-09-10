package model

import (
	"time"

	"gorm.io/gorm"
)

type App struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"uniqueIndex"`
	CreatedAt time.Time `json:"createdAt"`
}
