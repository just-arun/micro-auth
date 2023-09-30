package grpcclient

import (
	"context"
	"fmt"

	"github.com/just-arun/micro-auth/model"
	pb "github.com/just-arun/micro-session-proto"
)



type siteMapSession struct{}

func SiteMapSession() siteMapSession {
	return siteMapSession{}
}



func (st siteMapSession) SetServiceMap(client pb.SessionServiceClient, siteMapData []model.ServiceMap) error {
	fmt.Println("client streaming started")
	stream, err := client.SetServiceMap(context.Background())
	if err != nil {
		return err
	}

	for _, v := range siteMapData {
		err := stream.Send(&pb.ServiceMapData{
			Id:    uint64(v.ID),
			Key:   v.Key,
			Value: v.Key,
			Auth:  v.Auth,
		})
		if err != nil {
			return err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Println(res.Ok)
	return nil
}

func (st siteMapSession) GetServiceMap(client pb.SessionServiceClient, ) (result []model.ServiceMap, err error) {
	stream, err := client.GetServiceMap(context.Background(), &pb.NoPayload{})
	if err != nil {
		fmt.Println("ERR: ", err.Error())
		return
	}
	for {
		resp, err := stream.Recv()	
		if err != nil {
			return []model.ServiceMap{}, nil
		}
		result = append(result, model.ServiceMap{
			ID: uint(resp.Id),
			Key: resp.Key,
			Value: resp.Value,
			Auth: resp.Auth,
		})
	}
}

