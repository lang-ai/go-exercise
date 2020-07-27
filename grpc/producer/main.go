package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	fmt.Println("# STARTING GRPC Producer")
	port := ":50051"
	err := StartProducer(port)
	if err != nil {
		log.Fatalf("Failed to start producer: %v", err)
	}
}
