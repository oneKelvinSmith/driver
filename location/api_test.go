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

	var (
		server   *httptest.Server
		response *http.Response
		err      error
	)

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
		store := &Store{}
		driverID := DriverID(21)
		driverLocations := []DriverLocation{
			DriverLocation{
				DriverID: driverID,
				Location: Location{
					Latitude:  42.42,
					Longitude: 22.33,
					UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
				},
			},
			DriverLocation{
				DriverID: driverID,
				Location: Location{
					Latitude:  42.424,
					Longitude: 22.332,
					UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
				},
			},
		}

		var locations []Location

		BeforeEach(func() {
			store.ConnectDB(":6379")
			for _, driverLocation := range driverLocations {
				store.PushLocation(driverLocation)
			}
			api.ConnectStore(store)
		})

		AfterEach(func() {
			store.DeleteLocations(driverID)
			Expect(err).To(BeNil())
		})

		It("retuns a JSON object containing the driver's location for a given time period", func() {
			response, err = http.Get(server.URL + "/drivers/21/coordinates?minutes=5")

			decoder := json.NewDecoder(response.Body)
			err = decoder.Decode(&locations)

			Expect(locations).To(Equal([]Location{
				Location{
					Latitude:  42.424,
					Longitude: 22.332,
					UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
				},
				Location{
					Latitude:  42.42,
					Longitude: 22.33,
					UpdatedAt: "YYYY-MM-DDTHH:MM:SSZ",
				},
			}))
		})
	})
})
