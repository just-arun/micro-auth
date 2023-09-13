package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID       uint     `gorm:"primaryKey"`
	Name     string   `json:"name"`
	Accesses []Access `json:"access" gorm:"many2many:role_access;"`
}
