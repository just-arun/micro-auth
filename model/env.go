package model

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

type Env struct {
	DB   db      `mapstructure:"db"`
	Grpc grpc    `mapstructure:"grpc"`
	Nats natsEnv `mapstructure:"nats"`
	Rsa  rsaEnv  `mapstructure:"rsa"`
}
