package main

// Driver represents the zombie status of a driver
type Driver struct {
	ID     int  `json:"id"`
	Zombie bool `json:"zombie"`
}

func main() {
	api := API{}
	api.Serve(":3000")
}
