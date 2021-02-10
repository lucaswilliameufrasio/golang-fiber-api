package main

import (
	"net/http"
	_ "net/http/pprof"

	"fmt"
	"log"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config/environment"
)

func main() {
	server := config.App()
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	// app.Use(logger.New())
	if err := server.Listen(fmt.Sprintf(":%v", environment.Port)); err != nil {
		log.Panic(err)
	}
}
