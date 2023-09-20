package model

import "gorm.io/gorm"

type ServiceMap struct {
	gorm.Model
	ID    uint   `json:"id" gorm:"primaryKey"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
