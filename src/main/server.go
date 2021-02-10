package main

import (
	"fmt"
	"log"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config/environment"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := config.App()

	go func() {
		if err := server.Listen(fmt.Sprintf(":%v", environment.Port)); err != nil {
			log.Panic(err)
		}
	}()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)

	_ = <-channel
	fmt.Println("Gracefully shutting down...")
	_ = server.Shutdown()
	fmt.Println("Closing connections...")
	// Put tasks to close connections or whatever needed down here
	// gormhelpers.Disconnect()
}
