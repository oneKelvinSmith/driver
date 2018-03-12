package main_test

import (
	"encoding/json"

	"github.com/nsqio/go-nsq"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "driver/location"
)

var _ = Describe("Consumer", func() {
	consumer := Consumer{}
	store := Store{}

	var (
		err     error
		payload []byte
	)

	BeforeEach(func() {
		store.ConnectDB(":6379")
		consumer.ConnectStore(&store)
	})

	AfterEach(func() {
		store.DeleteLocations(42)
		Expect(err).To(BeNil())
	})

	Describe("UpdateDriverLocation", func() {
		It("registers a message consumer with NSQ that subscribes to a given topic and channel", func() {
			driverLocation := DriverLocation{
				DriverID: 42,
				Location: Location{
					Latitude:  48.8566,
					Longitude: 2.3522,
					UpdatedAt: "2018-03-12T02:09:52Z",
				},
			}

			messageID := [16]byte{42}
			payload, err = json.Marshal(driverLocation)

			message := nsq.NewMessage(messageID, payload)

			err = consumer.UpdateLocation(message)

			location := store.GetLastLocation(driverLocation.DriverID)

			Expect(location).To(Equal(Location{
				Latitude:  48.8566,
				Longitude: 2.3522,
				UpdatedAt: "2018-03-12T02:09:52Z",
			}))
		})
	})
})
