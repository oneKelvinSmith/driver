package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestZombie(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zombie Suite")
}
