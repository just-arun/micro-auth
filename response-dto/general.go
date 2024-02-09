package responsedto

import (
	"time"

	"github.com/just-arun/micro-auth/model"
)

type GeneralGetOne struct {
	ID                      uint                 `json:"id"`
	Name                    string               `json:"name"`
	CanLogin                bool                 `json:"canLogin"`
	CanRegister             bool                 `json:"canRegister"`
	HttpOnlyCookie          bool                 `json:"httpOnlyCookie"`
	AccessTokenExpiryTime   time.Duration        `json:"accessTokenExpireTime"`
	RefreshTokenExpiryTime  time.Duration        `json:"refreshTokenExpireTime"`
	OrganizationEmailDomain string               `json:"organizationEmailDomain"`
	TokenPlacement          model.TokenPlacement `json:"tokenPlacement"`
	UpdatedBy               []SlimUser           `json:"updatedBy"`
	UpdatedDescription      string               `json:"updatedDescription"`
}
