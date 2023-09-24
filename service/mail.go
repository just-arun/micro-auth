package service

import (
	"github.com/just-arun/micro-auth/model"
	"gorm.io/gorm"
)

type mail struct{}

func Mail() mail {
	return mail{}
}

func (st mail) SendOtp(db *gorm.DB, to, message string) error {
	return db.Save(&model.Mail{Message: message, To: to}).Error
}


