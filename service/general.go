package service

import (
	"github.com/just-arun/micro-auth/model"
	"gorm.io/gorm"
)

type general struct{}

func General() general {
	return general{}
}

func (r general) Create(db *gorm.DB, data *model.General) error {
	return db.Save(data).Error
}

func (r general) Get(db *gorm.DB) (data *model.General, err error) {
	err = db.First(&data).Error
	return
}

func (r general) Update(db *gorm.DB, id uint, data *model.General) (err error) {
	data.ID = id
	return db.Save(&data).Error
}
