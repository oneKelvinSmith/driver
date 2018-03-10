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

// Connect initialised the store and creates a redis pool.
func (s *Store) Connect() {
	s.pool = &redis.Pool{
		MaxIdle: 3,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
}

// GetConnection returns a redis connection.
func (s *Store) GetConnection() redis.Conn {
	return s.pool.Get()
}

// InsertLocation stores a driver's latest location in to redis.
func (s *Store) InsertLocation(u LocationUpdate) {
	key := "location:" + string(u.DriverID)
	value, err := json.Marshal(u.Location)
	handleStoreError(err)

	_, err = s.GetConnection().Do("SET", key, value)
	handleStoreError(err)
}

func handleStoreError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
