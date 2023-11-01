package service

import (
	"fmt"

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

func (u user) CreateMultiple(db *gorm.DB, users []*model.User) error {
	for i, _ := range users {
		users[i].Type = model.UserTypeUnVerify
	}
	tnx := db.Save(&users)
	if tnx.Error != nil {
		return tnx.Error
	}
	return nil
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
		First(&filter)
	if tnx.Error != nil {
		return nil, tnx.Error
	}
	return filter, nil
}

func (u user) GetMany(db *gorm.DB, filter *model.User) (data []model.User, err error) {
	fmt.Println("ROLES", 23)
	err = db.
		Model(&filter).
		Preload("Roles").
		Scan(&data).
		Error
	return
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

func (u user) UpdateRole(db *gorm.DB, userID uint, roles []model.Role) error {
	return db.Model(&model.User{ID: userID}).Association("Roles").Replace(roles)
}
