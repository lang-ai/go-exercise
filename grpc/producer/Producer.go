package main

import (
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ps struct{}

// message creates a random message
func (p *ps) message() string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GetMessage returns a single message
func (p *ps) GetMessage(_ context.Context, _ *Empty) (*Message, error) {
	m := p.message()
	log.Println("Produce message", m)
	return &Message{Payload: m}, nil
}

// StartProducer starts a new producer on a given port
func StartProducer(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterProducerServer(s, &ps{})
	log.Println("Producer started on port:", port)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
