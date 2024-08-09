package main

import (
	"context"
	"log"
	"time"

	"github.com/popeskul/awesome-messanger/services/profile/pb/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50054"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewProfileServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	updateResp, err := client.UpdateProfile(ctx, &proto.UpdateProfileRequest{
		UserId:    "123",
		Nickname:  "new_nickname",
		Bio:       "This is a new bio.",
		AvatarUrl: "http://example.com/avatar.jpg",
	})
	if err != nil {
		log.Fatalf("could not update profile: %v", err)
	}

	log.Printf("Update Profile Response: Success: %v, Message: %s", updateResp.GetSuccess(), updateResp.GetMessage())

	getResp, err := client.GetProfile(ctx, &proto.GetProfileRequest{
		UserId: "123",
	})
	if err != nil {
		log.Fatalf("could not get profile: %v", err)
	}

	log.Printf("Get Profile Response: UserId: %s, Nickname: %s, Bio: %s, AvatarUrl: %s",
		getResp.GetUserId(), getResp.GetNickname(), getResp.GetBio(), getResp.GetAvatarUrl())
}
