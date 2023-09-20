package grpcclient

import (
	"context"
	"time"

	pb "github.com/just-arun/micro-session-proto"
)

type userSession struct{}

func UserSession() userSession {
	return userSession{}
}

func (s userSession) SetUserSession(client pb.SessionServiceClient, userID uint, roles []string) (*pb.SetUserSessionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SetUserSession(ctx, &pb.UserSessionPayload{
		UserID: 143,
		Role:   roles,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
