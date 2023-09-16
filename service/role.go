package service

import (
	"github.com/just-arun/micro-auth/model"
	"gorm.io/gorm"
)

type role struct{}

func Role() role {
	return role{}
}

func (r role) GetNames(db *gorm.DB) (roles []model.Role, err error) {
	err = db.Find(&model.Role{}).Scan(&roles).Error
	return
}

func (r role) Add(db *gorm.DB, role model.Role) (err error) {
	err = db.Save(role).Error
	return
}

func (r role) GetOne(db *gorm.DB, id uint) (role *model.Role, err error) {
	err = db.Preload("Access").First(&model.Role{ID: id}).Scan(&role).Error
	return
}
