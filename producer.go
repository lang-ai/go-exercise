package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Producer produces new events sent to a channel
type Producer struct {
	id         int  // id to identify the producer
	waitMillis int  // millis to wait between production
	started    bool // whether the producer has started producing
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// message creates a random message
func (p *Producer) message() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Start makes the producer start publishing in the specified channel
func (p *Producer) Start(ch chan<- string) {
	fmt.Printf("Producer[%v]: Starting\n", p.id)
	if p.started {
		panic("producer already started!")
	}
	p.started = true
	go func() {
		for {
			msg := p.message()
			fmt.Printf("Producer[%v] produced message: %q\n", p.id, msg)
			ch <- msg
			time.Sleep(time.Duration(p.waitMillis) * time.Millisecond)
		}
	}()
}

var numProducers int

// NewProducer returns a new Producer
func NewProducer(waitMillis int) *Producer {
	numProducers++
	return &Producer{id: numProducers, waitMillis: waitMillis, started: false}
}
