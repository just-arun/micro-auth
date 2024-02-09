package model

import "gorm.io/gorm"

type Mail struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	To      string
	Message string
}


func (Mail) TableName() string {
	return "mailes"
}
