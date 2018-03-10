package main

// Location represents a geographical location.
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UpdatedAt string  `json:"updated_at"`
}

// DriverID is the driver's unique identificatin number.
type DriverID int

// DriverLocation represents an update to the geographical location of a driver.
type DriverLocation struct {
	DriverID DriverID `json:"driver_id"`
	Location Location
}

func main() {
	store := Store{}
	store.ConnectDB("redis:6379")

	consumer := Consumer{}
	consumer.ConnectStore(&store)
	consumer.ConnectBus("nsqlookupd:4161", "driver", "location")

	api := API{}
	api.ConnectStore(&store)
	api.Serve(":3000")
}
