package pubsub

import "github.com/nats-io/nats.go"

type pub struct {
	Con *nats.EncodedConn
}

func Publisher() *pub {
	return &pub{}
}

func (st pub) ChangeServiceMap(con *nats.EncodedConn, data interface{}) error {
	return con.Publish("change-service-map", data)
}
