package grpcclient

import (
	"context"
	"fmt"
	"time"

	pb "github.com/just-arun/micro-session-proto"
)

type userSession struct{}

func UserSession() userSession {
	return userSession{}
}

func (s userSession) SetUserSession(client pb.SessionServiceClient, userID uint, role string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	
	res, err := client.SetUserSession(ctx, &pb.UserSessionPayload{
		UserID: 143,
		Role:   "admin",
	})
	if err != nil {
		return err
	}

	fmt.Println("Access Token", res.Token)
	fmt.Println("Refresh Token", res.Refresh)

	return nil
}
