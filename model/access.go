package model

type Access struct {
	ID   uint   `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name"`
	Key  string `json:"key"`
}
