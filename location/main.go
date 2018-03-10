package main

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
