package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "driver/zombie"
)

var _ = Describe("Categoriser", func() {
	categoriser := Categoriser{}

	Describe("Categorise", func() {
		driver := &Driver{ID: DriverID(42), Zombie: false}

		Context("when theres is not enough location data has", func() {
			locations := []Location{
				Location{},
			}

			It("does NOT mark the driver as a zombie", func() {
				categoriser.Categorise(driver, locations)

				Expect(driver.Zombie).To(BeFalse())
			})
		})

		Context("when driver has not moved", func() {
			locations := []Location{
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
			}

			It("marks the driver as a zombie", func() {
				categoriser.Categorise(driver, locations)

				Expect(driver.Zombie).To(BeTrue())
			})
		})

		Context("when driver is active", func() {
			locations := []Location{
				Location{
					Latitude:  53.354,
					Longitude: 76.911,
				},
				Location{
					Latitude:  53.352,
					Longitude: 76.909,
				},
				Location{
					Latitude:  53.35,
					Longitude: 76.90,
				},
			}

			It("does NOT mark the driver as a zombie", func() {
				categoriser.Categorise(driver, locations)

				Expect(driver.Zombie).To(BeFalse())
			})
		})
	})
})
