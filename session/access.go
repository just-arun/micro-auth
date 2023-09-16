package session

import (
	"github.com/just-arun/micro-auth/model"
	"gorm.io/gorm"
)

type access struct{}

func Access() access {
	return access{}
}

func (r access) AddOne(db *gorm.DB, access model.Access) error {
	return db.Save(&access).Error
}

func (r access) GetAll(db *gorm.DB) (accesses []model.Access, err error) {
	err = db.Find(&model.Access{}).Scan(&accesses).Error
	return
}

func (r access) DeleteOne(db *gorm.DB, id uint) (err error) {
	err = db.Delete(&model.Access{ID: id}).Error
	return
}
