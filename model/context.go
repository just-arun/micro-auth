package model

import (
	pb "github.com/just-arun/micro-session-proto"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type HandlerCtx struct {
	Env            *Env
	DB             *gorm.DB
	GrpcClient     *pb.SessionServiceClient
	NatsConnection *nats.EncodedConn
}
