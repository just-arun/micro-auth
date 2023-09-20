package service

import (
	"github.com/just-arun/micro-auth/model"
	"gorm.io/gorm"
)

type serviceMap struct{}

func ServiceMap() serviceMap {
	return serviceMap{}
}

func (st serviceMap) Add(db *gorm.DB, data *model.ServiceMap) error {
	return db.Save(&data).Error
}

func (st serviceMap) GetOne(db *gorm.DB, id uint) (data *model.ServiceMap, err error) {
	err = db.Model(&model.ServiceMap{ID: id}).Scan(&data).Error
	return
}

func (st serviceMap) GetMany(db *gorm.DB) (data []model.ServiceMap, err error) {
	err = db.Find(&data).Error
	return
}

func (st serviceMap) UpdateOne(db *gorm.DB, id uint, data *model.ServiceMap) (err error) {
	data.ID = id
	return db.Save(&data).Error
}

func (st serviceMap) DeleteOne(db *gorm.DB, id uint) (err error) {
	return db.Delete(&model.ServiceMap{ID: id}).Error
}



