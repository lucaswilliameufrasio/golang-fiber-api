package main

import (
	"fmt"
	"log"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config/environment"
)

func main() {
	server := config.App()

	if err := server.Listen(fmt.Sprintf(":%v", environment.Port)); err != nil {
		log.Panic(err)
	}
}
