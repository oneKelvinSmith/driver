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
		conn = store.Pool.Get()
	})

	AfterEach(func() {
		store.DeleteLocations(driverID)
		_ = conn.Close()
		Expect(err).To(BeNil())
	})

	Describe("PushLocation", func() {
		driverLocation := DriverLocation{
			DriverID: driverID,
			Location: Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: "2018-03-12T02:09:41Z",
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
				UpdatedAt: "2018-03-12T02:09:41Z",
			}))
		})
	})

	Describe("GetLastLocation", func() {
		driverLocation := DriverLocation{
			DriverID: driverID,
			Location: Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: "2018-03-12T02:09:41Z",
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
				UpdatedAt: "2018-03-12T02:09:41Z",
			}))
		})

		It("returns an empty location if there is no data in redis", func() {
			location := store.GetLastLocation(0)

			Expect(location).To(Equal(Location{}))
		})
	})

	Describe("GetLocations", func() {
		driverLocations := []DriverLocation{
			DriverLocation{
				DriverID: driverID,
				Location: Location{
					Latitude:  51.51,
					Longitude: 18.18,
					UpdatedAt: "2018-03-12T02:09:50Z",
				},
			},
			DriverLocation{
				DriverID: driverID,
				Location: Location{
					Latitude:  15.15,
					Longitude: 81.81,
					UpdatedAt: "2018-03-12T02:09:52Z",
				},
			},
			// DriverLocation{
			//	DriverID: driverID,
			//	Location: Location{
			//		Latitude:  15.15,
			//		Longitude: 81.81,
			//		UpdatedAt: "2018-03-12T03:09:52Z",
			//	},
			// },
		}

		BeforeEach(func() {
			for _, driverLocation := range driverLocations {
				store.PushLocation(driverLocation)
			}
		})

		It("retrieves all the locations for the last 5 minutes for a given DriverID from redis", func() {
			locations := store.GetLocations(driverID)

			Expect(locations).To(Equal([]Location{
				Location{
					Latitude:  15.15,
					Longitude: 81.81,
					UpdatedAt: "2018-03-12T02:09:52Z",
				},
				Location{
					Latitude:  51.51,
					Longitude: 18.18,
					UpdatedAt: "2018-03-12T02:09:50Z",
				},
			}))
		})

		It("returns an empty slice if there is no data in redis", func() {
			locations := store.GetLocations(0)

			Expect(locations).To(BeEmpty())
		})
	})
})
