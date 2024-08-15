package main

import (
	"context"
	"log"
	"time"

	"github.com/popeskul/awesome-messanger/services/message/pkg/api/message"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50060"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := message.NewMessageServiceClient(conn)

	testGetMessages(client)
	testSendMessage(client)
}

func testGetMessages(client message.MessageServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &message.GetMessagesRequest{
		ChatId: "chat1", // Replace with a valid chat ID
	}

	resp, err := client.GetMessages(ctx, req)
	if err != nil {
		log.Fatalf("could not get messages: %v", err)
	}

	for _, msg := range resp.GetMessages() {
		log.Printf("Message: %s from %s", msg.GetContent(), msg.GetSenderId())
	}
}

func testSendMessage(client message.MessageServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &message.SendMessageRequest{
		SenderId:    "user1",
		RecipientId: "user2",
		Content:     "Hello there!",
	}

	resp, err := client.SendMessage(ctx, req)
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}

	if resp.GetSuccess() {
		log.Println("Message sent successfully!")
	} else {
		log.Println("Failed to send message.")
	}
}
