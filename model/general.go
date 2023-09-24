package model

import (
	"time"

	"gorm.io/gorm"
)

type General struct {
	gorm.Model
	ID                      uint   `json:"id" gorm:"primaryKey"`
	Name                    string `json:"name" gorm:"uniqueIndex;default:app"`
	CanLogin                bool
	CanRegister             bool
	HttpOnlyCookie          bool
	AccessTokenExpiryTime   time.Duration
	RefreshTokenExpiryTime  time.Duration
	OrganizationEmailDomain string
}





