package main


import (
	"log"
	"mbatimel/WB_Tech-L2/tree/main/develop/dev11/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	server.SetupHandlers()
	server.Up()
}