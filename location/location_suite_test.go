package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLocation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Location Suite")
}
