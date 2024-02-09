package model

type ServiceMap struct {
	ID      uint   `json:"id,omitempty" gorm:"primaryKey"`
	Key     string `json:"key,omitempty" gorm:"uniqueIndex" mapestructure:"key"`
	Value   string `json:"value,omitempty" mapestructure:"value"`
	Auth    bool   `json:"auth" gorm:"default:false" mapestructure:"auth"`
	Default bool   `json:"default" mapestructure:"default"`
}

func (ServiceMap) TableName() string {
	return "service_maps"
}
