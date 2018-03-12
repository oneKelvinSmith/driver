package main

func main() {
	categoriser := &Categoriser{}

	api := API{Categoriser: categoriser}
	api.Serve(":3002")
}
