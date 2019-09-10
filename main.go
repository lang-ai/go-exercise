package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("HOLA!")

	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	messagesBus := make(chan string)
	totalProducers := 2
	totalConsumers := 1
	for i := 0; i < totalProducers; i++ {
		p := NewProducer(500 + seededRand.Intn(500))
		p.Start(messagesBus)
	}
	for i := 0; i < totalConsumers; i++ {
		c := NewConsumer()
		c.Start(messagesBus)
	}

	time.Sleep(10 * time.Second)
}
