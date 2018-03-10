package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "driver/location"

	"github.com/garyburd/redigo/redis"
)

var _ = Describe("Store", func() {
	var connection redis.Conn
	var err error

	Describe("InsertLocation", func() {
		BeforeEach(func() {
			connection, err = redis.Dial("tcp", ":6379")
		})

		AfterEach(func() {
			connection.Close()
			Expect(err).To(BeNil())
		})

		It("inserts the location into redis", func() {
			connection.Do("SET", "location", "coords")

			var result interface{}
			result, err = connection.Do("GET", "location")

			Expect(string(result.([]byte))).To(Equal("coords"))
		})
	})

})
