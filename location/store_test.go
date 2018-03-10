package main_test

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "driver/location"
)

var _ = Describe("Store", func() {
	store := &Store{}

	var (
		connection redis.Conn
		err        error
	)

	Describe("InsertLocation", func() {
		BeforeEach(func() {
			store.Connect(":6379")
		})

		AfterEach(func() {
			err = connection.Close()
			Expect(err).To(BeNil())
		})

		It("inserts the location into redis", func() {
			driverLocation := DriverLocation{
				DriverID: 42,
				Location: Location{
					Latitude:  48.8566,
					Longitude: 2.3522,
					UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
				},
			}

			store.InsertLocation(driverLocation)

			var (
				key      string
				value    interface{}
				location Location
			)

			connection = store.GetConnection()
			key = "location:" + string(driverLocation.DriverID)
			value, err = connection.Do("GET", key)
			err = json.Unmarshal(value.([]byte), &location)

			expectedLocation := Location{
				Latitude:  48.8566,
				Longitude: 2.3522,
				UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
			}

			Expect(location).To(Equal(expectedLocation))
		})
	})

})
