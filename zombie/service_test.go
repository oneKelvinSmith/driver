package main_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "nsq/zombie"
)

var _ = Describe("Service", func() {
	service := &Service{}
	service.Initialise()

	var server *httptest.Server

	BeforeEach(func() {
		server = httptest.NewServer(service.Router)
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

		It("displays a welcome message with the handler", func() {
			request := httptest.NewRequest(http.MethodGet, "/", nil)
			recorder := httptest.NewRecorder()

			service.GetHome(recorder, request)

			response := recorder.Result()

			body, err := ioutil.ReadAll(response.Body)
			Expect(err).To(BeNil())
			Expect(response.Body.Close()).To(BeNil())

			Expect(string(body)).To(Equal("Welcome to Zombie"))
		})
	})
})
