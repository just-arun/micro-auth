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
	Name                    string         `json:"name" gorm:"uniqueIndex;default:app"`
	CanLogin                bool           `json:"canLogin"`
	CanRegister             bool           `json:"canRegister"`
	HttpOnlyCookie          bool           `json:"httpOnlyCookie"`
	AccessTokenExpiryTime   time.Duration  `json:"accessTokenExpireTime"`
	RefreshTokenExpiryTime  time.Duration  `json:"refreshTokenExpireTime"`
	OrganizationEmailDomain string         `json:"organizationEmailDomain"`
	TokenPlacement          TokenPlacement `json:"tokenPlacement" gorm:"type:token_placement;default:header"`
}
