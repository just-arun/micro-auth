package api

import (
	"encoding/json"
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

func apiV1(e *echo.Echo, g *echo.Group, environment, port, context string) {

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

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	fmt.Println(string(data))

	// 	var data []model.ServiceMap
	// 	aclData := `package acl

	// type ACL string

	// const (
	// `
	// 	for _, v := range util.GetPath(e) {
	// 		da := model.ServiceMap{
	// 			Key:   v.Key,
	// 			Value: strings.ReplaceAll(v.Value, ".", " "),
	// 			Auth:  true,
	// 		}
	// 		aclData += fmt.Sprintf(`   ACL%v ACL = "%v"
	//     `,
	// 			strings.ReplaceAll(strings.ReplaceAll(v.Value, "auth.", ""), ".", ""),
	// 			v.Key,
	// 		)
	// 		data = append(data, da)
	// 	}
	// 	aclData += `)
	// `
	// 	os.WriteFile("acl/acl.go", []byte(aclData), 0644)

	// 	for _, v := range data {
	// 		pDb.Save(&v)
	// 	}

	byteData, err := service.ServiceMap().GetMany(ctx.DB)
	if err != nil {
		panic(err)
	}
	_ = pubsub.Publisher().ChangeServiceMap(natsConnection, byteData)

	serverPort := fmt.Sprintf(":%v", port)
	e.Logger.Fatal(
		e.Start(serverPort),
	)
}
