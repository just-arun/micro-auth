package pubsub

import (
	"encoding/json"
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
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Println("++++++++++++")
	fmt.Println(byteData)
	fmt.Println("++++++++++++")
	fmt.Println(string(byteData))
	fmt.Println("++++++++++++")
	return con.Publish("change-service-map", data)
}
