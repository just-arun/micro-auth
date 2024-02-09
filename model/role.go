package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID                  uint     `json:"id" gorm:"primaryKey"`
	Name                string   `json:"name" gorm:"uniqueIndex,required"`
	Accesses            []Access `json:"accesses" gorm:"many2many:role_access;"`
	IsScheduledToDelete bool     `json:"isScheduledToDelete" gorm:"default:false"`
}


func (Role) TableName() string {
	return "roles"
}


