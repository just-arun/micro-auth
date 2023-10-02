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
	err = db.
		// Preload(clause.Associations).
		Preload("Accesses").
		Find(&roles).Error
	// Model(&model.Role{}).Scan(&roles).Error
	return
}

func (r role) Add(db *gorm.DB, role *model.Role) (err error) {
	err = db.Save(&role).Error
	return
}

func (r role) AddAccess(db *gorm.DB, id uint, access *model.Access) (err error) {
	role := &model.Role{ID: id}
	tnx := db.First(&role)
	if tnx.Error != nil {
		return err
	}
	role.Accesses = append(role.Accesses, *access)
	if tnx.Save(role).Error != nil {
		return
	}
	return
}

func (r role) GetOne(db *gorm.DB, id uint) (role *model.Role, err error) {
	err = db.Preload("Access").First(&model.Role{ID: id}).Scan(&role).Error
	return
}
