package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

// Consumer represents an abstraction over NSQ.
type Consumer struct {
	consumer *nsq.Consumer
	store    *Store
}

// ConnectBus registers the consumer with a NSQ daemon via the lookup daemon
// to subscribe to a given topic and channel.
func (c *Consumer) ConnectBus(lookupHost string, topic string, channel string) {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	handleConsumerError(err)

	consumer.ChangeMaxInFlight(100)
	consumer.AddHandler(nsq.HandlerFunc(c.UpdateLocation))

	err = consumer.ConnectToNSQLookupd(lookupHost)
	handleConsumerError(err)

	c.consumer = consumer
}

// ConnectStore attaches the consumer to the redis backed store.
func (c *Consumer) ConnectStore(s *Store) {
	c.store = s
}

// UpdateLocation is a message handler that stores driver location updates.
func (c *Consumer) UpdateLocation(m *nsq.Message) error {
	driverLocation := DriverLocation{
		Location: Location{
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
	}
	err := json.Unmarshal(m.Body, &driverLocation)

	if err != nil {
		return err
	}

	c.store.PushLocation(driverLocation)

	return nil
}

func handleConsumerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
