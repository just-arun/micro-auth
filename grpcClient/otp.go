package grpcclient

import (
	"context"
	"time"

	"github.com/just-arun/micro-auth/model"
	pb "github.com/just-arun/micro-session-proto"
)

type otpSession struct{}

func OTPSession() otpSession {
	return otpSession{}
}

func (s otpSession) SetOTP(client pb.SessionServiceClient, OTP string, userID uint, key model.OTPKey) (*pb.OkResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.SetOTP(ctx, &pb.OTPPayload{
		Otp:    OTP,
		UserID: uint64(userID),
		Key:    string(key),
	})
	if err != nil {
		return nil, err
	}
	return &pb.OkResponse{Ok: response.Ok}, nil
}

func (s otpSession) GetOTP(client pb.SessionServiceClient, OTP string, userID uint, key model.OTPKey) (*pb.OkResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.GetOTP(ctx, &pb.OTPPayload{
		Otp:    OTP,
		UserID: uint64(userID),
		Key:    string(key),
	})
	if err != nil {
		return nil, err
	}
	return &pb.OkResponse{Ok: response.Ok}, nil
}
