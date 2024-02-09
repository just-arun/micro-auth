package model

type RoleAccess struct {
	AccessID uint   `json:"access_id,omitempty"`
	RoleID   uint   `json:"role_id" gorm:"primaryKey"`
	RoleName string `json:"role_name,omitempty" gorm:"role_name"`
}

func (RoleAccess) TableName() string {
	return "role_access"
}
