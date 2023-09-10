package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	UserName  string    `json:"userName"`
	Password  string    `json:"password"`
	Apps      []App     `json:"apps" gorm:"type:integer[];foreignKey:ID;references:ID"`
	CreatedAt time.Time `json:"createdAt"`
}
