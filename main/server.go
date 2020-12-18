package main

import (
	"log"
	config "lucaswilliameufrasio/golang-fiber-api/main/config"
)

func main() {
	server := config.App()
	// app.Use(logger.New())
	log.Fatal(server.Listen(":7979"))
}
