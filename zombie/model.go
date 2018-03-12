package main

// Location represents a geographical location.
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// DriverID represents the unique identifier of the driver.
type DriverID int

// Driver represents the zombie status of a driver.
type Driver struct {
	ID     DriverID `json:"id"`
	Zombie bool     `json:"zombie"`
}
