/*
Main client backend

*/
package main

import (
	"context"
	"log"

	"github.com/AntonyIS/golang-grpc-app-1/chat"
	"google.golang.org/grpc"
)

func main() {
	// Create a client connection instance
	var conn *grpc.ClientConn
	// Try to connect to the gRPC server on port 9000
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	// Check for errors if connection was unsuccesful
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Close the connection before closing the main go routine(main func)
	defer conn.Close()

	// Connect with the chat service in the server, using the connection from the server
	c := chat.NewChatServiceClient(conn)

	// Create a message payload to send to the chat service
	message := chat.Message{
		Body: "Hello from client",
	}
	// Send the message to the chat service
	response, err := c.Sayhallo(context.Background(), &message)
	// Handle errors with the connection was unsuccesful
	if err != nil {
		log.Fatalf("Error calling the Sayhallo %s", err)
	}
	// Log the response from the client on the console
	log.Println(response.Body)
}
