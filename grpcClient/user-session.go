package grpcclient

import (
	"context"
	"fmt"
	"net/http"
	"strings"
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
		UserID: uint64(userID),
		Roles:  roles,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (st userSession) HaveAccess(client pb.SessionServiceClient, r *http.Request, routeAccessReferenceKey string) (*pb.HaveAccessResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	roles := r.Header.Get("x-roles")

	if len(roles) == 0 {
		return nil, fmt.Errorf("unauthorized")
	}

	res, err := client.HaveAccess(ctx, &pb.HaveAccessParam{Roles: strings.Split(roles, ","), AccessSlug: routeAccessReferenceKey})
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	if !res.Access {
		return nil, fmt.Errorf("unauthorized")
	}

	return res, nil
}
