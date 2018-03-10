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
	key := "driver:" + string(driverID) + ":location"

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
		conn.Do("DEL", key)
		err = conn.Close()
		Expect(err).To(BeNil())
	})

	Describe("PushLocation", func() {
		driverLocation := DriverLocation{
			DriverID: driverID,
			Location: Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
			},
		}

		It("inserts the location into redis", func() {
			store.PushLocation(driverLocation)

			var values [][]byte
			values, err = redis.ByteSlices(
				conn.Do("LRANGE", key, "0", "0"),
			)
			err = json.Unmarshal(values[0], &location)

			Expect(location).To(Equal(Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
			}))
		})
	})

	Describe("GetLastLocation", func() {
		driverLocation := DriverLocation{
			DriverID: driverID,
			Location: Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
			},
		}

		BeforeEach(func() {
			store.PushLocation(driverLocation)
		})

		It("retrieves the location for a given DriverID from redis", func() {
			location := store.GetLastLocation(driverID)

			Expect(location).To(Equal(Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
			}))
		})

		It("returns an empty location if there is no data in redis", func() {
			location := store.GetLastLocation(0)

			Expect(location).To(Equal(Location{}))
		})
	})
})
