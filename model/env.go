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
	CanLogin                bool          `mapstructure:"can_login"`
	CanRegister             bool          `mapstructure:"can_register"`
	HttpOnlyCookie          bool          `mapstructure:"http_only_cookie"`
	AccessTokenExpiryTime   time.Duration `mapstructure:"access_token_expiry_time"`
	RefreshTOkenExpiryTime  time.Duration `mapstructure:"refresh_token_expiry_time"`
	OrganizationEmailDomain string        `mapstructure:"organization_email_domain"`
}

type userSeed struct {
	UserName string `mapsructure:"userName"`
	Email    string `mapsructure:"email"`
	Password string `mapsructure:"password"`
}

type Env struct {
	DB         db           `mapstructure:"db"`
	Grpc       grpc         `mapstructure:"grpc"`
	MailGrpc   grpc         `mapstructure:"mailGrpc"`
	Nats       natsEnv      `mapstructure:"nats"`
	Rsa        rsaEnv       `mapstructure:"rsa"`
	General    general      `mapstructure:"general"`
	Admin      userSeed     `mapstructure:"admin"`
	ServiceMap []ServiceMap `mapstructure:"serviceMap"`
}
