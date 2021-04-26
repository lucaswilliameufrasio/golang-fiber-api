package events_test

import (
	"lucaswilliameufrasio/golang-fiber-api/src/main/config"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func SUTWebsocket() *fiber.App {
	return config.App(fiber.Config{
		DisableStartupMessage: true,
	})
}

func TestWebsocket(t *testing.T) {
	app := SUTWebsocket()
	go app.Listen(":7777")
	defer app.Shutdown()

	u := "ws://localhost:7777/api/ws"

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Errorf("Got an error: %v", err)
	}
	defer ws.Close()

	// Send message to server, read response and check to see if it's what we expect.
	if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		t.Errorf("Got an error: %v", err)
	}
	_, message, err := ws.ReadMessage()
	if err != nil {
		t.Errorf("Got an error: %v", err)
	}
	assert.Equal(t, string(message), "hello")
}
