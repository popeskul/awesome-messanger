package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/auth"
)

const (
	address = "localhost:50050"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := auth.NewAuthServiceClient(conn)

	loginResp, err := c.Login(context.Background(), &auth.LoginRequest{
		Username: "testuser",
		Password: "testpassword",
	})
	if err != nil {
		log.Fatalf("VerifyCredentials failed: %v", err)
	}

	if loginResp.GetToken() == "" {
		log.Fatal("VerifyCredentials response token is empty")
	}

	log.Printf("VerifyCredentials Response: %s", loginResp.GetToken())

	registerResp, err := c.Register(context.Background(), &auth.RegisterRequest{
		Username: "newuser",
		Password: "newpassword",
	})
	if err != nil {
		log.Fatalf("Register failed: %v", err)
	}
	log.Printf("Register Response User: %s", registerResp.GetUser())
	log.Printf("Register Response Token: %s", registerResp.GetToken())

	refreshResp, err := c.Refresh(context.Background(), &auth.RefreshRequest{
		OldToken: loginResp.GetToken(),
	})
	if err != nil {
		log.Fatalf("Refresh failed: %v", err)
	}

	log.Printf("Refresh Response: %s", refreshResp.GetNewToken())

	if _, err := c.Logout(context.Background(), &auth.LogoutRequest{
		Token: loginResp.GetToken(),
	}); err != nil {
		log.Fatalf("Logout failed: %v", err)
	}

	log.Println("Logout success")
}
