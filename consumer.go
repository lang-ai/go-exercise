package main

import "fmt"

// Consumer consumes new events sent to a channel
type Consumer struct {
	id      int  // id to identify the consumer
	started bool // whether the consumer has started consuming
}

// Start makes the consumer start publishing in the specified channel
func (c *Consumer) Start(ch <-chan string) {
	fmt.Printf("Consumer[%v]: Starting\n", c.id)
	if c.started {
		panic("consumer already started!")
	}
	c.started = true
	go func() {
		for {
			msg := <-ch
			fmt.Printf("Consumer[%v]: produced message: %q\n", c.id, msg)
		}
	}()
}

var numConsumers int

// NewConsumer returns a new Consumer
func NewConsumer() *Consumer {
	numConsumers++
	return &Consumer{id: numConsumers}
}
