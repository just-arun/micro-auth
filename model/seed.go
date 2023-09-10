package model

type userSeed struct {
	UserName string `mapsructure:"userName"`
	Email    string `mapsructure:"email"`
	Password string `mapsructure:"password"`
}

type Seed struct {
	Admin userSeed `mapstructure:"admin"`
}
