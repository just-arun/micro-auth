package model

type ServiceMap struct {
	ID      uint   `json:"id,omitempty" gorm:"primaryKey"`
	Key     string `json:"key,omitempty" gorm:"uniqueIndex"`
	Value   string `json:"value,omitempty"`
	Auth    bool   `json:"auth" gorm:"default:false"`
	Default bool   `json:"default"`
}
