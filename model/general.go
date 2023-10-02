package model

import (
	"time"

	"gorm.io/gorm"
)

type General struct {
	gorm.Model
	ID                      uint          `json:"id"`
	Name                    string        `json:"name" gorm:"uniqueIndex;default:app"`
	CanLogin                bool          `json:"canLogin"`
	CanRegister             bool          `json:"canRegister"`
	HttpOnlyCookie          bool          `json:"httpOnlyCookie"`
	AccessTokenExpiryTime   time.Duration `json:"accessTokenExpireTime"`
	RefreshTokenExpiryTime  time.Duration `json:"refreshTokenExpireTime"`
	OrganizationEmailDomain string        `json:"OrganizationEmailDomain"`
}
