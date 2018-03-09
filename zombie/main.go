package main

func main() {
	service := Service{}
	service.Initialise()
	service.Start(":3000")
}
