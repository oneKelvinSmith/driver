package main_test

import (
	"encoding/json"
	"time"

	"github.com/garyburd/redigo/redis"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "driver/location"
)

func currentTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

func pastTimestamp() string {
	return time.Now().Add(-10 * time.Minute).Format(time.RFC3339)
}

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
		updatedAt := currentTimestamp()
		driverLocation := DriverLocation{
			DriverID: driverID,
			Location: Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: updatedAt,
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
				UpdatedAt: updatedAt,
			}))
		})
	})

	Describe("GetLastLocation", func() {
		updatedAt := currentTimestamp()
		driverLocation := DriverLocation{
			DriverID: driverID,
			Location: Location{
				Latitude:  48.48,
				Longitude: 3.33,
				UpdatedAt: updatedAt,
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
				UpdatedAt: updatedAt,
			}))
		})

		It("returns an empty location if there is no data in redis", func() {
			location := store.GetLastLocation(0)

			Expect(location).To(Equal(Location{}))
		})
	})

	Describe("GetLocations", func() {
		updatedAt := currentTimestamp()
		oldUpdatedAt := pastTimestamp()

		driverLocations := []DriverLocation{
			DriverLocation{
				DriverID: driverID,
				Location: Location{
					Latitude:  111.111,
					Longitude: 222.222,
					UpdatedAt: updatedAt,
				},
			},
			DriverLocation{
				DriverID: driverID,
				Location: Location{
					Latitude:  11.11,
					Longitude: 22.22,
					UpdatedAt: updatedAt,
				},
			},
			DriverLocation{
				DriverID: driverID,
				Location: Location{
					Latitude:  15.15,
					Longitude: 81.81,
					UpdatedAt: oldUpdatedAt,
				},
			},
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
					Latitude:  111.111,
					Longitude: 222.222,
					UpdatedAt: updatedAt,
				},
				Location{
					Latitude:  11.11,
					Longitude: 22.22,
					UpdatedAt: updatedAt,
				},
			}))
		})

		It("returns an empty slice if there is no data in redis", func() {
			locations := store.GetLocations(0)

			Expect(locations).To(BeEmpty())
		})
	})
})
