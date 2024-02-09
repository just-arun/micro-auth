package model

import (
	pb "github.com/just-arun/micro-session-proto"
	pbMailing "github.com/just-arun/micro-session-proto/mailing"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type HandlerCtx struct {
	Env            *Env
	DB             *gorm.DB
	GrpcClient     *pb.SessionServiceClient
	MailGrpcClient *pbMailing.MailingNotificationServiceClient
	NatsConnection *nats.EncodedConn
}
