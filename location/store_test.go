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
	driverID := 42

	var (
		key      string
		value    interface{}
		location Location
		conn     redis.Conn
		err      error
	)

	BeforeEach(func() {
		store.ConnectDB(":6379")
		conn = store.Connect()
	})

	AfterEach(func() {
		err = conn.Close()
		Expect(err).To(BeNil())
	})

	Describe("GetLocation", func() {
		It("retrieves the location for a given DriverID from redis", func() {
			key = "location:" + string(driverID)
			value, err = conn.Do("SET", key, "some_value")

			location := store.GetLocation(driverID)
			Expect(location).To(Equal([]byte("some_value")))
		})
	})

	Describe("SetLocation", func() {
		It("inserts the location into redis", func() {
			driverLocation := DriverLocation{
				DriverID: driverID,
				Location: Location{
					Latitude:  48.8566,
					Longitude: 2.3522,
					UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
				},
			}

			store.SetLocation(driverLocation)

			key = "location:" + string(driverLocation.DriverID)
			value, err = conn.Do("GET", key)
			err = json.Unmarshal(value.([]byte), &location)

			Expect(location).To(Equal(Location{
				Latitude:  48.8566,
				Longitude: 2.3522,
				UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
			}))
		})
	})

})
