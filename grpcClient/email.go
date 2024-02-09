package grpcclient

import (
	"context"
	"fmt"
	"time"

	pb "github.com/just-arun/micro-session-proto/mailing"
)

type emailSession struct{}

func EmailSession() emailSession {
	return emailSession{}
}

func (s emailSession) SetMail(mailClient pb.MailingNotificationServiceClient, payload *pb.MailSendTwoFactorAuthOtpPayload) (*pb.OkResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	fmt.Println(payload.FromName)
	fmt.Println(payload.FromEmail)
	fmt.Println(payload.ToName)
	fmt.Println(payload.ToEmail)
	fmt.Println(payload.Link)
	fmt.Println(payload.MailType)
	fmt.Println(payload.Otp)

	response, err := mailClient.SendMail(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &pb.OkResponse{Ok: response.Ok}, nil
}
