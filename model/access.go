package model

import "gorm.io/gorm"

type Access struct {
	gorm.Model
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Key  string `json:"key"`
}














