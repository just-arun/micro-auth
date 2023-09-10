package model

type db struct {
	Uri string `mapstructure:"uri"`
}

type Env struct {
	DB db `mapstructure:"db"`
}




