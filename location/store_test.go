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

	var connection redis.Conn
	var err error

	Describe("InsertLocation", func() {
		BeforeEach(func() {
			store.Connect()
		})

		AfterEach(func() {
			err = connection.Close()
			Expect(err).To(BeNil())
		})

		// Reference JSON payload
		// {
		//	"driver_id": 42,
		//	"location": {
		//		"latitude": 48.8566,
		//		"longitude": 2.3522,
		//		"updated_at": "YYYY-MM-DDTHH:MM:SSZ"
		//	}
		// }

		It("inserts the location into redis", func() {
			locationUpdate := LocationUpdate{
				DriverID: 42,
				Location: Location{
					Latitude:  48.8566,
					Longitude: 2.3522,
					UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
				},
			}

			store.InsertLocation(locationUpdate)

			var key string
			var value interface{}
			var location Location

			connection = store.GetConnection()
			key = "location:" + string(locationUpdate.DriverID)
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
