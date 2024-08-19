package main

import (
	"context"
	"log"
	"time"

	"github.com/popeskul/awesome-messanger/services/message/pb/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052" // Replace with the actual address of your gRPC server
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client for the MessageService.
	client := proto.NewMessageServiceClient(conn)

	// Test the client methods.
	testGetMessages(client)
	testSendMessage(client)
}

func testGetMessages(client proto.MessageServiceClient) {
	// Set up a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Prepare the request.
	req := &proto.GetMessagesRequest{
		ChatId: "chat1", // Replace with a valid chat ID
	}

	// Call the GetMessages method.
	resp, err := client.GetMessages(ctx, req)
	if err != nil {
		log.Fatalf("could not get messages: %v", err)
	}

	// Log the response.
	for _, msg := range resp.GetMessages() {
		log.Printf("Message: %s from %s", msg.GetContent(), msg.GetSenderId())
	}
}

func testSendMessage(client proto.MessageServiceClient) {
	// Set up a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Prepare the request.
	req := &proto.SendMessageRequest{
		SenderId:    "user1",        // Replace with the sender's ID
		RecipientId: "user2",        // Replace with the recipient's ID
		Content:     "Hello there!", // Replace with the actual message content
	}

	// Call the SendMessage method.
	resp, err := client.SendMessage(ctx, req)
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}

	// Log the response.
	if resp.GetSuccess() {
		log.Println("Message sent successfully!")
	} else {
		log.Println("Failed to send message.")
	}
}
