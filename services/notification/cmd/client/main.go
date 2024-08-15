package main

import (
	"context"
	"log"
	"time"

	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50070"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := notification.NewNotificationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := client.SendNotification(ctx, &notification.SendNotificationRequest{
		RecipientId: "user123",
		Message:     "Hello, this is a test notification!",
	})
	if err != nil {
		log.Fatalf("could not send notification: %v", err)
	}

	log.Printf("Notification sent successfully: %v", resp)
}
