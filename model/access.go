package model

import "gorm.io/gorm"

type Access struct {
	gorm.Model
	ID   uint
	Name string
	Key  string
}
