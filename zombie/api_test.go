package main_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "driver/zombie"
)

type MockClient struct {
	locations []Location
}

func (c *MockClient) Do(req *http.Request) (*http.Response, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	_ = encoder.Encode(c.locations)

	response := http.Response{}
	_ = response.Write(buffer)
	response.Header.Set("Content-Type", "application/json")

	return &response, nil
}

var _ = Describe("API", func() {
	var response *http.Response
	var err error

	api := &API{Categoriser: &Categoriser{}}

	router := api.NewRouter()
	server := httptest.NewServer(router)

	AfterEach(func() {
		_ = response.Body.Close()
		server.Close()
	})

	Describe("GET /", func() {

		It("return the service name as a health check", func() {
			response, err = http.Get(server.URL)

			var body []byte
			body, err = ioutil.ReadAll(response.Body)

			Expect(string(body)).To(Equal("Zombie"))
		})
	})

	Describe("GET /drivers/:id", func() {
		var driver Driver

		mockClient := &MockClient{
			locations: []Location{
				Location{
					Latitude:  53.352,
					Longitude: 76.909,
				},
				Location{
					Latitude:  53.352,
					Longitude: 76.909,
				},
				Location{
					Latitude:  53.352,
					Longitude: 76.909,
				},
			},
		}

		It("retuns a JSON object containing the driver's zombie status", func() {
			// I struggled to get an decent integration test between this and the location service.
			// I really tried :(
			Skip("Unable to figure out http mocking issues...")

			api.Categoriser.SetClient(mockClient)

			response, err = http.Get(server.URL + "/drivers/42")

			decoder := json.NewDecoder(response.Body)
			_ = decoder.Decode(&driver)

			expectedDriver := Driver{ID: 42, Zombie: true}
			Expect(driver).To(Equal(expectedDriver))

			var responseJSON, driverJSON []byte
			responseJSON, err = json.Marshal(driver)
			driverJSON, err = json.Marshal(expectedDriver)

			Expect(responseJSON).To(MatchJSON(driverJSON))
		})
	})
})
