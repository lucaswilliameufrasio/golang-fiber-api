package main

import (
	"fmt"
	"log"
	config "lucaswilliameufrasio/golang-fiber-api/src/main/config"
)

func main() {
	server := config.App()

	// app.Use(logger.New())
	if err := server.Listen(fmt.Sprintf(":%v", config.Port)); err != nil {
		log.Panic(err)
	}
}
