package requestdto

type CreateServiceMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Auth  bool   `json:"bool"`
}
