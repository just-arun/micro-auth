package service

import (
	"github.com/just-arun/micro-auth/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type user struct{}

func User() user {
	return user{}
}

func (u user) CreateOne(db *gorm.DB, user *model.User) (uint, error) {
	user.Type = model.UserTypeUnVerify
	tnx := db.Save(&user)
	if tnx.Error != nil {
		return 0, tnx.Error
	}
	return user.ID, nil
}

func (u user) UpdateVerify(db *gorm.DB, userID uint) error {
	user := &model.User{ID: userID}
	tnx := db.First(&user)
	if tnx.Error != nil {
		return tnx.Error
	}
	user.Type = model.UserTypeVerified
	return tnx.Save(&user).Error
}

func (u user) GetOne(db *gorm.DB, filter *model.User) (*model.User, error) {
	tnx := db.
		Model(&model.User{}).
		Preload(clause.Associations).
		// Preload("Apps").
		First(&filter)
	if tnx.Error != nil {
		return nil, tnx.Error
	}
	return filter, nil
}

func (u user) AddApp(db *gorm.DB, userID uint, apps []model.App) error {
	user := &model.User{}
	tnx := db.First(&user, "id = ?", userID).Scan(&user)
	if tnx.Error != nil {
		return tnx.Error
	}
	user.Apps = append(user.Apps, apps...)
	tnx = tnx.Save(&user)
	if tnx.Error != nil {
		return tnx.Error
	}
	return nil
}
