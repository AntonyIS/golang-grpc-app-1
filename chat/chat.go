package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Sayhallo(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message from client: %s", message.Body)
	return &Message{Body: "Hello from the server"}, nil
}
