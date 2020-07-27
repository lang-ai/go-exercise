package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("# STARTING GRPC Consumer")
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewProducerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMessage(ctx, &Empty{})
	if err != nil {
		log.Fatalf("could not get the message: %v", err)
	}
	log.Printf("Message: %s", r.GetPayload())
}
