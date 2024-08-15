package main

import (
	"context"
	"log"
	"time"

	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/profile"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50080"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := profile.NewProfileServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err = client.UpdateProfile(ctx, &profile.UpdateProfileRequest{
		UserId:    "1",
		Nickname:  "popeskul",
		Bio:       "I am a software engineer",
		AvatarUrl: "https://avatars.githubusercontent.com/u/12345678",
	})
}
