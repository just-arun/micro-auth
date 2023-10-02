package model

import "time"

type db struct {
	Uri string `mapstructure:"uri"`
}

type grpc struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type natsEnv struct {
	Token string `mapstructure:"token"`
}

type rsaEnv struct {
	PrivateKey string `mapstructure:"privateKey"`
}

type general struct {
	CanLogin                bool          `mapstructure:"canLogin"`
	CanRegister             bool          `mapstructure:"canRegister"`
	HttpOnlyCookie          bool          `mapstructure:"httpOnlyCookie"`
	AccessTokenExpiryTime   time.Duration `mapstructure:"accessTokenExpiryTime"`
	RefreshTOkenExpiryTime  time.Duration `mapstructure:"refreshTOkenExpiryTime"`
	OrganizationEmailDomain string        `mapstructure:"organizationEmailDomain"`
}

type userSeed struct {
	UserName string `mapsructure:"userName"`
	Email    string `mapsructure:"email"`
	Password string `mapsructure:"password"`
}

type Env struct {
	DB         db         `mapstructure:"db"`
	Grpc       grpc       `mapstructure:"grpc"`
	Nats       natsEnv    `mapstructure:"nats"`
	Rsa        rsaEnv     `mapstructure:"rsa"`
	General    general    `mapstructure:"general"`
	Admin      userSeed   `mapstructure:"admin"`
	ServiceMap []ServiceMap `mapstructure:"serviceMap"`
}
