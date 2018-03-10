package main

import (
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

// Consumer represents an abstraction over NSQ.
type Consumer struct {
	consumer *nsq.Consumer
	store    *Store
}

// ConnectBus registers the consumer with a NSQ daemon via the lookup daemon
// to subscribe to a given topic and channel.
func (c *Consumer) ConnectBus(port string, topic string, channel string) {
	consumer, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	handleConsumerError(err)

	consumer.AddHandler(nsq.HandlerFunc(c.UpdateLocation))

	err = consumer.ConnectToNSQD(port)
	handleConsumerError(err)

	c.consumer = consumer
}

// ConnectStore attaches the consumer to the redis backed store.
func (c *Consumer) ConnectStore(s *Store) {
	c.store = s
}

// UpdateLocation is a message handler that stores driver location updates.
func (c *Consumer) UpdateLocation(m *nsq.Message) error {
	driverLocation := DriverLocation{}
	err := json.Unmarshal(m.Body, &driverLocation)

	if err == nil {
		c.store.PushLocation(driverLocation)
	}

	return err
}

func handleConsumerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
