package routes_test

import (
	"fmt"
	"io/ioutil"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func SUTStatus() *fiber.App {
	return config.App(fiber.Config{
		DisableStartupMessage: true,
	})
}

func TestMainRoute(t *testing.T) {
	app := SUTStatus()

	// http.Request
	req := httptest.NewRequest("GET", "http://localhost:7777/api/", nil)

	// http.Response
	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, `{"data":"Hello, World 👋!"}`, string(body))

	// Do something with results:
	if resp.StatusCode != 200 {
		fmt.Println(string(body)) // => Hello, World 👋!

		t.Errorf("got %d, want %d", resp.StatusCode, 200)
	}
}
