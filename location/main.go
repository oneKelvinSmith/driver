package main

// Location represents a geographical location.
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UpdatedAt string  `json:"updated_at"`
}

// LocationUpdate represents an update to the geographical location of a driver.
type LocationUpdate struct {
	DriverID int `json:"driver_id"`
	Location Location
}

func main() {
	api := API{}
	api.Serve(":3000")
	store := Store{}
	store.Connect(":6379")
}
