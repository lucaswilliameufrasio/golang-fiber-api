package events

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

// WebsocketHandler is a function to handle websocket messages
func WebsocketHandler(c *websocket.Conn) {
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)
		err = c.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
