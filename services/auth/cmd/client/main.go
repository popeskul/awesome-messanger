package main

import (
	"context"
	"log"

	"github.com/popeskul/awesome-messanger/services/auth/pb/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewAuthServiceClient(conn)

	loginResp, err := c.Login(context.Background(), &proto.LoginRequest{
		Username: "testuser",
		Password: "testpassword",
	})
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	if loginResp.GetToken() == "" {
		log.Fatal("Login response token is empty")
	}

	log.Printf("Login Response: %s", loginResp.GetToken())

	registerResp, err := c.Register(context.Background(), &proto.RegisterRequest{
		Username: "newuser",
		Password: "newpassword",
	})
	if err != nil {
		log.Fatalf("Register failed: %v", err)
	}
	log.Printf("Register Response: %s", registerResp.GetMessage())

	refreshResp, err := c.Refresh(context.Background(), &proto.RefreshRequest{
		OldToken: loginResp.GetToken(),
	})
	if err != nil {
		log.Fatalf("Refresh failed: %v", err)
	}

	log.Printf("Refresh Response: %s", refreshResp.GetNewToken())

	logoutResp, err := c.Logout(context.Background(), &proto.LogoutRequest{
		Token: loginResp.GetToken(),
	})
	if err != nil {
		log.Fatalf("Logout failed: %v", err)
	}

	log.Printf("Logout Response: %s", logoutResp.GetMessage())
}
