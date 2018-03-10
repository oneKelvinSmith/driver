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

	driverID := DriverID(42)

	var (
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

	Describe("SetLocation", func() {
		driverLocation := DriverLocation{
			DriverID: driverID,
			Location: Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
			},
		}

		var (
			key   string
			value interface{}
		)

		It("inserts the location into redis", func() {
			store.SetLocation(driverLocation)

			key = "location:" + string(driverID)
			value, err = conn.Do("GET", key)
			err = json.Unmarshal(value.([]byte), &location)

			Expect(location).To(Equal(Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
			}))
		})
	})

	Describe("GetLocation", func() {
		driverLocation := DriverLocation{
			DriverID: driverID,
			Location: Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
			},
		}

		BeforeEach(func() {
			store.SetLocation(driverLocation)
		})

		It("retrieves the location for a given DriverID from redis", func() {
			location := store.GetLocation(driverID)

			Expect(location).To(Equal(Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
			}))
		})
	})
})
