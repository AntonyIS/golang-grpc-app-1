/*
Main server backend
*/
package main

import (
	"log"
	"net"

	"example.com/app_1/chat"
	"google.golang.org/grpc"
)

func main() {
	// Create a server that listns on port 9000
	listen, err := net.Listen("tcp", ":9000")

	// Check for errors if the server cant server on port 9000
	if err != nil {
		log.Fatalf("Failed to connect to server on port 9000 : %v", err)
	}
	log.Println("Server running on port 9000")

	// Initialize the server structure from the chat package
	s := chat.Server{}

	// Initialize the gRPC server
	grpcServer := grpc.NewServer()

	// Register the chat service, taking grpcServer and the structure variable
	chat.RegisterChatServiceServer(grpcServer, &s)

	// Serve the grpcserver that listens on the same port as the main server : port 9000
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000 : %v", err)
	}
}
