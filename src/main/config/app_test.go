// +test
package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func SUT() *fiber.App {
	return App(fiber.Config{
		DisableStartupMessage: true,
	})
}

func TestWebsocket(t *testing.T) {
	app := SUT()
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

func TestMainRoute(t *testing.T) {
	app := SUT()

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

func TestCorrectEmailAndPassword(t *testing.T) {
	app := SUT()

	type Response struct {
		Token string `json:"token"`
		Email string `json:"email"`
	}

	// http.Request
	httpRequest := []byte(`{"email":"user1@example.com","password":"pass"}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v1/login", bytes.NewBuffer(httpRequest))
	req.Header.Set("Content-Type", "application/json")

	// http.Response
	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	result := &Response{}
	err := json.Unmarshal(body, result)

	if err != nil {
		fmt.Println(err)
	}

	userEmail := result.Email

	assert.Equal(t, "user1@example.com", userEmail)
}

func TestWrongEmail(t *testing.T) {
	app := SUT()

	// http.Request
	httpRequest := []byte(`{"email":"userunknown@example.com","password":"pass"}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v1/login", bytes.NewBuffer(httpRequest))
	req.Header.Set("Content-Type", "application/json")

	// http.Response
	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	expectedResponse := `{"error":"Unauthorized"}`

	assert.Equal(t, expectedResponse, string(body))
}

func TestWrongPassword(t *testing.T) {
	app := SUT()

	// http.Request
	httpRequest := []byte(`{"email":"user1@example.com","password":"not_valid"}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v1/login", bytes.NewBuffer(httpRequest))
	req.Header.Set("Content-Type", "application/json")

	// http.Response
	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	expectedResponse := `{"error":"Unauthorized"}`

	assert.Equal(t, expectedResponse, string(body))
}

func TestProtectedRoute(t *testing.T) {
	app := SUT()

	type LoginResponse struct {
		Token string `json:"token"`
		Email string `json:"email"`
	}

	// http.Request
	httpRequest := []byte(`{"email":"user1@example.com","password":"pass"}`)

	reqLogin := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v1/login", bytes.NewBuffer(httpRequest))
	reqLogin.Header.Set("Content-Type", "application/json")

	respWithToken, _ := app.Test(reqLogin)

	body, _ := ioutil.ReadAll(respWithToken.Body)

	resultOfLogin := &LoginResponse{}
	err := json.Unmarshal(body, resultOfLogin)

	if err != nil {
		fmt.Println(err)
	}

	accessToken := resultOfLogin.Token

	mainRequest := httptest.NewRequest(http.MethodGet, "http://localhost:7777/api/v1/protected", nil)
	mainRequest.Header.Set("Content-Type", "application/json")
	mainRequest.Header.Set("Authorization", "Bearer "+accessToken)

	responseActual, _ := app.Test(mainRequest)
	responseBody, _ := ioutil.ReadAll(responseActual.Body)
	expectedResponse := `{"data":"Hello, Dude 👋! Your ID is 1"}`

	assert.Equal(t, expectedResponse, string(responseBody))
}

func TestProtectedRouteError(t *testing.T) {
	app := SUT()

	req := httptest.NewRequest(http.MethodGet, "http://localhost:7777/api/v1/protected", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	expectedResponse := `{"error":"Access Denied"}`

	assert.Equal(t, expectedResponse, string(body))
}
