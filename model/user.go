package model

import (
	"time"

	"github.com/just-arun/micro-auth/util"
	"gorm.io/gorm"
)

type UserType string

const (
	UserTypeUnVerify UserType = "UN_VERIFIED"
	UserTypeVerified UserType = "VERIFIED"
	UserTypeDisabled UserType = "DISABLED"
	UserTypeDeleted  UserType = "DELETED"
)

func RegisterUserType() string {
	return `CREATE TYPE user_type AS ENUM (
		'UN_VERIFIED',
		'VERIFIED',
		'DISABLED',
		'DELETED'
	);`
}

// func (ct *UserType) Scan(value interface{}) error {
// 	*ct = UserType(value.([]byte))
// 	return nil
// }

// func (ct UserType) Value() (driver.Value, error) {
// 	return string(ct), nil
// }

type User struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	UserName  string    `json:"userName"`
	Password  string    `json:"password,omitempty"`
	Type      UserType  `json:"type" gorm:"type:user_type;default:UN_VERIFIED"`
	Roles     []Role    `json:"roles" gorm:"many2many:user_role;foreignKey:ID"`
	Apps      []App     `json:"apps" gorm:"many2many:user_app;foreignKey:ID"`
	Profile   Profile   `json:"profile" gorm:"foreignKey:ID;references:ID"`
	CreatedAt time.Time `json:"createdAt"`
}

type SlimUser struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Email    string   `json:"email" gorm:"uniqueIndex"`
	UserName string   `json:"userName"`
	Type     UserType `json:"type"`
	Roles    []struct {
		Name string `json:"name"`
	} `json:"roles"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *User) HashPassword(tx *gorm.DB) (err error) {
	if len(u.Password) > 20 {
		return nil
	}
	hash, err := util.Password().Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

// gorm hook
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	return u.HashPassword(tx)
}

// gorm hook
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return u.HashPassword(tx)
}
