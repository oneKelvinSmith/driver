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

// SetLocation stores a driver's latest location in to redis.
func (s *Store) SetLocation(d DriverLocation) {
	key := "location:" + string(d.DriverID)
	value, err := json.Marshal(d.Location)
	handleStoreError(err)

	_, err = s.Connect().Do("SET", key, value)
	handleStoreError(err)
}

// GetLocation stores a driver's latest location in to redis.
func (s *Store) GetLocation(d DriverID) Location {
	key := "location:" + string(d)
	value, err := s.Connect().Do("GET", key)
	handleStoreError(err)

	var location Location
	err = json.Unmarshal(value.([]byte), &location)
	handleStoreError(err)

	return location
}

func handleStoreError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
