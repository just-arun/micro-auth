package service

import (
	"github.com/just-arun/micro-auth/model"
	"gorm.io/gorm"
)


type app struct {}

func App() app {
	return app{}
}

func (a app) CreateOne(db *gorm.DB, ap *model.App) (uint, error) {
	tnx := db.Create(ap)
	if tnx.Error != nil {
		return 0, tnx.Error
	}
	return ap.ID, nil
}

func (a app) GetOne(db *gorm.DB, filter *model.App) (ap *model.User, err error) {
	tnx := db.
		Find(filter).
		Scan(ap)
	if tnx.Error != nil {
		return nil, tnx.Error
	}
	return ap, nil
}


