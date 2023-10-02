package service

import (
	"fmt"

	"github.com/just-arun/micro-auth/acl"
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

func (r role) AddMultipleAccess(db *gorm.DB, id uint, access []model.Access) (err error) {
	role := &model.Role{ID: id}
	tnx := db.First(&role)
	if tnx.Error != nil {
		return err
	}
	role.Accesses = append(role.Accesses, access...)
	if tnx.Save(role).Error != nil {
		return
	}
	return
}

func (r role) GetOne(db *gorm.DB, id uint) (role *model.Role, err error) {
	err = db.Preload("Access").First(&model.Role{ID: id}).Scan(&role).Error
	return
}

func (r role) PopulateBasicRole(db *gorm.DB) (err error) {
	accesses, err := Access().GetMany(db, []string{
		string(acl.ACLAuthLogin),
		string(acl.ACLAuthGetPublicKey),
	})

	if err != nil {
		return err
	}

	fmt.Println(accesses)

	var role model.Role
	role.Name = "basic"
	role.Accesses = accesses

	fmt.Println(role)
	db.Delete(&model.Role{})
	_ = db.Save(&role).Error
	_ = r.AddMultipleAccess(db, role.ID, accesses)
	return
}
