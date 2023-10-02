package model

type Access struct {
	ID   uint   `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex"`
	Key  string `json:"key" gorm:"uniqueIndex"`
}
