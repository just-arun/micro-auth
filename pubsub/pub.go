package pubsub

import (
	"fmt"

	"github.com/just-arun/micro-auth/model"
	"github.com/nats-io/nats.go"
)

type pub struct {
	Con *nats.EncodedConn
}

func Publisher() *pub {
	return &pub{}
}

func (st pub) ChangeServiceMap(con *nats.EncodedConn, data []model.ServiceMap) error {
	fmt.Println(data)
	return con.Publish("change-service-map", data)
}
