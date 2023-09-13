package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	ID        uint `json:"id" gorm:"primaryKey"`
	City      string
	State     string
	Country   string
	CreatedAt time.Time
}
