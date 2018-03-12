package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Store is used to insert and retrieve driver location state.
type Store struct {
	protocol string
	port     string

	Pool *redis.Pool
}

// ConnectDB initialised the store and creates a redis pool.
func (s *Store) ConnectDB(port string) {
	s.port = port
	s.Pool = s.newPool()
	s.cleanup()
}

func (s *Store) newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", s.port)
			if err != nil {
				return nil, err
			}
			return conn, err
		},
	}
}

func (s *Store) cleanup() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, syscall.SIGTERM)
	signal.Notify(channel, syscall.SIGKILL)
	go func() {
		<-channel
		_ = s.Pool.Close()
		os.Exit(0)
	}()
}

// PushLocation stores a driver's latest location in a redis list.
func (s *Store) PushLocation(d DriverLocation) {
	value, err := json.Marshal(d.Location)
	handleStoreError(err)

	conn := s.Pool.Get()

	_, err = conn.Do("LPUSH", key(d.DriverID), value)
	handleStoreError(err)

	err = conn.Close()
	handleStoreError(err)
}

// GetLastLocation stores a driver's latest location in to redis.
func (s *Store) GetLastLocation(id DriverID) Location {
	conn := s.Pool.Get()

	values, err := redis.ByteSlices(conn.Do("LRANGE", key(id), 0, 0))
	handleStoreError(err)

	err = conn.Close()
	handleStoreError(err)

	if len(values) > 0 {
		var location Location
		err = json.Unmarshal(values[0], &location)
		handleStoreError(err)

		return location
	}

	return Location{}
}

// GetLocations stores a driver's latest location in to redis.
func (s *Store) GetLocations(id DriverID) []Location {
	conn := s.Pool.Get()
	values, err := redis.ByteSlices(conn.Do("LRANGE", key(id), 0, -1))
	handleStoreError(err)

	err = conn.Close()
	handleStoreError(err)

	locations := []Location{}
	var location Location
	for i := len(values) - 1; i >= 0; i-- {
		err = json.Unmarshal(values[i], &location)
		handleStoreError(err)

		if fiveMinutesAgo() < updatedAt(location) {
			locations = append(locations, location)
		} else {
			return locations
		}
	}

	return locations
}

// DeleteLocations removes all location data for a given driver.
func (s *Store) DeleteLocations(id DriverID) {
	conn := s.Pool.Get()
	_, err := conn.Do("DEL", key(id))
	handleStoreError(err)
}

func fiveMinutesAgo() int64 {
	t := time.Now().Add(-5 * time.Minute)

	return t.Unix()
}

func updatedAt(l Location) int64 {
	t, err := time.Parse(time.RFC3339, l.UpdatedAt)
	handleStoreError(err)

	return t.Unix()
}

func key(id DriverID) string {
	return "driver:" + string(id) + ":location"
}

func handleStoreError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
