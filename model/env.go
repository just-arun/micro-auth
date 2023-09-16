package model

type db struct {
	Uri string `mapstructure:"uri"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Env struct {
	DB             db    `mapstructure:"db"`
	UserSession    Redis `mapstructure:"userSession"`
	GeneralSession Redis `mapstructure:"generalSession"`
}
