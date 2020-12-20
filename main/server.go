package main

import (
	"log"
	config "lucaswilliameufrasio/golang-fiber-api/main/config"
	"os"
)

func main() {
	server := config.App()
	PORT := os.Getenv("PORT")
	var address = ":7979"
	if PORT != "" {
		address = ":" + PORT
	}
	// app.Use(logger.New())
	log.Fatal(server.Listen(address))
}
