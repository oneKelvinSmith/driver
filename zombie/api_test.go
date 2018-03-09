package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "nsq/zombie"
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
		It("return the service name as a health check", func() {
			response, err = http.Get(server.URL)

			var body []byte
			body, err = ioutil.ReadAll(response.Body)

			Expect(string(body)).To(Equal("Zombie"))
		})
	})

	Describe("GET /drivers/:id", func() {
		var driver Driver

		It("retuns a JSON object containing the driver's zombie status", func() {
			response, err = http.Get(server.URL + "/drivers/42")

			decoder := json.NewDecoder(response.Body)
			err = decoder.Decode(&driver)

			expectedDriver := Driver{ID: 42, Zombie: true}
			Expect(driver).To(Equal(expectedDriver))

			var responseJSON, driverJSON []byte
			responseJSON, err = json.Marshal(driver)
			driverJSON, err = json.Marshal(expectedDriver)

			Expect(responseJSON).To(MatchJSON(driverJSON))
		})
	})
})
