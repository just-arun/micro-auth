package service

import (
	"github.com/just-arun/micro-auth/model"
	"gorm.io/gorm"
)

type user struct{}

func User() user {
	return user{}
}

func (u user) CreateOne(db *gorm.DB, user *model.User) (uint, error) {
	tnx := db.Model(&model.User{}).Save(&user)
	if tnx.Error != nil {
		return 0, tnx.Error
	}
	return user.ID, nil
}

func (u user) GetOne(db *gorm.DB, filter *model.User) (user *model.User, err error) {
	tnx := db.
		Find(filter).
		Scan(&user)
	if tnx.Error != nil {
		return nil, tnx.Error
	}
	return user, nil
}

func (u user) AddApp(db *gorm.DB, userID uint, app uint) error {
	if err := db.Model(&model.User{ID: userID}).
	Update("apps", []model.App{{ID: app}}).Error; err != nil {
		return err
	}
	return nil
}