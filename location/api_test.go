package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "driver/location"
)

var _ = Describe("API", func() {
	api := &API{}
	router := api.NewRouter()

	var server *httptest.Server
	var response *http.Response
	var err error

	BeforeEach(func() {
		server = httptest.NewServer(router)
	})

	AfterEach(func() {
		server.Close()
		Expect(response.Body.Close()).To(BeNil())
		Expect(err).To(BeNil())
	})

	Describe("GET /", func() {
		It("returns the service name as a health check", func() {
			response, err = http.Get(server.URL)

			var body []byte
			body, err = ioutil.ReadAll(response.Body)

			Expect(string(body)).To(Equal("Location"))
		})
	})

	Describe("GET /drivers/:id/coordinates", func() {
		var locations []Location

		It("retuns a JSON object containing the driver's location for a given time period", func() {
			response, err = http.Get(server.URL + "/drivers/21/coordinates?minutes=5")

			decoder := json.NewDecoder(response.Body)
			err = decoder.Decode(&locations)

			// Reference JSON response
			//
			// [
			//	{
			//		"latitude": 42,
			//		"longitude": 2.3,
			//		"updated_at": "YYYY-MM-DDTHH:MM:SSZ"
			//	},
			//	{
			//		"latitude": 42.1,
			//		"longitude": 2.32,
			//		"updated_at": "YYYY-MM-DDTHH:MM:SSZ"
			//	}
			// ]

			expectedLocations := []Location{
				Location{
					Latitude:  42,
					Longitude: 2.3,
					UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
				},
				Location{Latitude: 42.1,
					Longitude: 2.32,
					UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
				},
			}

			Expect(locations).To(Equal(expectedLocations))
		})
	})
})
