package main

import (
	"encoding/json"
	"log"

	"github.com/garyburd/redigo/redis"
)

// Store is used to insert and retrieve driver location state.
type Store struct {
	pool *redis.Pool
}

// ConnectDB initialised the store and creates a redis pool.
func (s *Store) ConnectDB(port string) {
	s.pool = &redis.Pool{
		MaxIdle: 3,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", port)
		},
	}
}

// Connect returns a redis connection.
func (s *Store) Connect() redis.Conn {
	return s.pool.Get()
}

// PushLocation stores a driver's latest location in a redis list.
func (s *Store) PushLocation(d DriverLocation) {
	value, err := json.Marshal(d.Location)
	handleStoreError(err)

	_, err = s.Connect().Do("LPUSH", key(d.DriverID), value)
	handleStoreError(err)
}

// GetLastLocation stores a driver's latest location in to redis.
func (s *Store) GetLastLocation(id DriverID) Location {
	values, err := redis.ByteSlices(
		s.Connect().Do("LRANGE", key(id), "0", "0"),
	)
	handleStoreError(err)

	if len(values) > 0 {
		var location Location
		err = json.Unmarshal(values[0], &location)
		handleStoreError(err)

		return location
	}

	return Location{}
}

func key(id DriverID) string {
	return "driver:" + string(id) + ":location"
}

func handleStoreError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
