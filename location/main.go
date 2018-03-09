package main

// Location represents the geographical location of a driver
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UpdatedAt string  `json:"updated_at"`
}

func main() {
	api := API{}
	api.Serve(":3000")
}
