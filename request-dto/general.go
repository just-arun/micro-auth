package requestdto

import (
	"time"

	"github.com/just-arun/micro-auth/model"
)

type UpdateGeneralAdminPayload struct {
	Name                    string `json:"name"`
	CanLogin                bool   `json:"canLogin"`
	CanRegister             bool   `json:"canRegister"`
	OrganizationEmailDomain string `json:"organizationEmailDomain"`
	UpdatedDescription      string `json:"updatedDescription"`
}

type UpdateGeneralServicePayload struct {
	HttpOnlyCookie          bool                 `json:"httpOnlyCookie"`
	AccessTokenExpiryTime   time.Duration        `json:"accessTokenExpireTime"`
	RefreshTokenExpiryTime  time.Duration        `json:"refreshTokenExpireTime"`
	TokenPlacement          model.TokenPlacement `json:"tokenPlacement"`
	UpdatedDescription      string               `json:"updatedDescription"`
}
