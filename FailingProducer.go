package main

import (
	"fmt"
	"time"
)

// FProducer produces new events sent to a channel, it might fail!
type FProducer struct {
	id         int  // id to identify the failing_producer
	waitMillis int  // millis to wait between production
	started    bool // whether the failing_producer has started producing
	failed     bool // whether the failing_producer has failed
}

// message creates a random message
func (p *FProducer) message() string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Start makes the producer start publishing in the specified channel
func (p *FProducer) Start(ch chan<- string) {
	fmt.Printf("FProducer[%v]: Starting\n", p.id)
	if p.started {
		panic("producer already started!")
	}
	p.started = true
	go func() {
		for {
			if seededRand.Float32() > 0.8 {
				fmt.Printf("FProducer[%v]: FAILED\n", p.id)
				p.failed = true
				break
			}
			msg := p.message()
			fmt.Printf("FProducer[%v] produced message: %q\n", p.id, msg)
			ch <- msg
			time.Sleep(time.Duration(p.waitMillis) * time.Millisecond)
		}
	}()
}

// NewFProducer returns a new FProducer
func NewFProducer(waitMillis int) *FProducer {
	numProducers++
	return &FProducer{id: numProducers, waitMillis: waitMillis, started: false, failed: false}
}
