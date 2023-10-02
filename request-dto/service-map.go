package requestdto

type CreateServiceMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Auth  bool   `json:"bool"`
}

type ServiceMapUpdate struct {
	ID    uint   `json:"id"`
	Key   string `json:"key,omitempty" gorm:"uniqueIndex" mapestructure:"key"`
	Value string `json:"value,omitempty" mapestructure:"value"`
	Auth  bool   `json:"auth" gorm:"default:false" mapestructure:"auth"`
}
