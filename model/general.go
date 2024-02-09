package model

import (
	"time"

	"gorm.io/gorm"
)

type TokenPlacement string

const (
	TokenPlacementHeader TokenPlacement = "header"
	TokenPlacementCookie TokenPlacement = "cookie"
)

func RegisterTokenPlacement() string {
	return `CREATE TYPE token_placement AS ENUM (
		'header',
		'cookie'
	);`
}

type General struct {
	gorm.Model
	ID                      uint           `json:"id"`
	Name                    string         `json:"name" gorm:"default:app"`
	CanLogin                bool           `json:"canLogin"`
	CanRegister             bool           `json:"canRegister"`
	HttpOnlyCookie          bool           `json:"httpOnlyCookie"`
	AccessTokenExpiryTime   time.Duration  `json:"accessTokenExpireTime"`
	RefreshTokenExpiryTime  time.Duration  `json:"refreshTokenExpireTime"`
	OrganizationEmailDomain string         `json:"organizationEmailDomain"`
	TokenPlacement          TokenPlacement `json:"tokenPlacement" gorm:"type:token_placement;default:header"`
	UpdatedBy               []User         `json:"updatedBy,omitempty" gorm:"many2many:general_updated_user;"`
	Author                  *GeneralAuthor `json:"author" gorm:"-:all"`
	UpdatedDescription      string         `json:"updatedDescription"`
	Active                  bool           `json:"active" gorm:"default:true"`
}

func (General) TableName() string {
	return "generals"
}

type GeneralAuthor struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"userName"`
}
