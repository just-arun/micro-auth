package grpcclient

import (
	"context"
	"fmt"

	"github.com/just-arun/micro-auth/model"
	pb "github.com/just-arun/micro-session-proto"
)

type role struct{}

func Role() role {
	return role{}
}

func (st role) SetRole(client pb.SessionServiceClient, role *model.Role) error {
	fmt.Println("setting role in session...")
	access := []*pb.AccessObject{}
	for _, v := range role.Accesses {
		access = append(access, &pb.AccessObject{
			Id:  uint64(v.ID),
			Key: v.Key,
		})
	}
	res, err := client.SetRole(context.Background(), &pb.RoleObject{
		Id:     uint64(role.ID),
		Name:   role.Name,
		Access: access,
	})
	if err != nil {
		return err
	}
	if res.Ok {
		return nil
	}
	return fmt.Errorf("error updating while updating")
}
