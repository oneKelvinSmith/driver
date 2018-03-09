package main_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "nsq/zombie"
)

var _ = Describe("API", func() {
	var api *API
	var server *httptest.Server

	BeforeEach(func() {
		api = &API{}
		server = httptest.NewServer(api.NewRouter())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("GET /", func() {
		It("displays a welcome message", func() {
			response, err := http.Get(server.URL)
			Expect(err).To(BeNil())

			body, err := ioutil.ReadAll(response.Body)
			Expect(err).To(BeNil())
			Expect(response.Body.Close()).To(BeNil())

			Expect(string(body)).To(Equal("Welcome to Zombie"))
		})
	})
})
