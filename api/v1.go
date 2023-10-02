package api

import (
	"fmt"

	"github.com/just-arun/micro-auth/boot"
	"github.com/just-arun/micro-auth/connections"
	"github.com/just-arun/micro-auth/model"
	"github.com/just-arun/micro-auth/pubsub"
	"github.com/just-arun/micro-auth/routes"
	"github.com/just-arun/micro-auth/service"
	"github.com/just-arun/micro-auth/util"
	pb "github.com/just-arun/micro-session-proto"
	"github.com/labstack/echo/v4"
)

func apiV1(e *echo.Echo, g *echo.Group, environment, port, context string, noServer ...bool) {

	env := &model.Env{}
	n := fmt.Sprintf(".env.%v", environment)
	util.GetEnv(n, context, &env)
	ctx := &model.HandlerCtx{}
	ctx.Env = env
	pDb := boot.PostgresDB(env.DB.Uri)
	ctx.DB = pDb

	conn := boot.NewGrpcConnection(env.Grpc.Host, env.Grpc.Port)

	client := pb.NewSessionServiceClient(conn)

	ctx.GrpcClient = &client

	natsConnection := boot.NatsConnection(env.Nats.Token)

	ctx.NatsConnection = natsConnection

	connections.HandlerCtx = ctx

	v1 := g.Group("/v1")
	routes.General(v1, ctx)
	routes.User(v1, ctx)
	routes.Role(v1, ctx)
	routes.Auth(v1, ctx)
	routes.Access(v1, ctx)
	routes.Role(v1, ctx)
	routes.ServiceMap(v1, ctx)

	byteData, err := service.ServiceMap().GetMany(ctx.DB)
	if err != nil {
		panic(err)
	}
	_ = pubsub.Publisher().ChangeServiceMap(natsConnection, byteData)

	serverPort := fmt.Sprintf(":%v", port)

	// check if server is required
	if len(noServer) == 1 {
		if noServer[0] {
			service.Access().GetSitemapAcl(e, pDb)
			service.Role().PopulateBasicRole(pDb)
			return
		}
	}

	e.Logger.Fatal(
		e.Start(serverPort),
	)
}
