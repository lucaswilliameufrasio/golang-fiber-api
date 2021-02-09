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
	"github.com/stretchr/testify/assert"
)

func SUT() *fiber.App {
	return App()
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

func TestProtectedRoute(t *testing.T) {
	app := SUT()

	type Response struct {
		Token string `json:"token"`
		Email string `json:"email"`
	}

	// http.Request
	httpRequest := []byte(`{"email":"user1@example.com","password":"pass"}`)

	reqLogin := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v1/login", bytes.NewBuffer(httpRequest))
	reqLogin.Header.Set("Content-Type", "application/json")

	respWithToken, _ := app.Test(reqLogin)

	body, _ := ioutil.ReadAll(respWithToken.Body)

	result := &Response{}
	err := json.Unmarshal(body, result)

	if err != nil {
		fmt.Println(err)
	}

	accessToken := result.Token

	mainRequest := httptest.NewRequest(http.MethodGet, "http://localhost:7777/api/v1/protected", nil)
	mainRequest.Header.Set("Content-Type", "application/json")
	mainRequest.Header.Set("Authorization", "Bearer "+accessToken)

	responseActual, _ := app.Test(mainRequest)
	responseBody, _ := ioutil.ReadAll(responseActual.Body)
	expectedResponse := `{"data":"Hello, Dude 👋! Your ID is 1"}`

	assert.Equal(t, expectedResponse, string(responseBody))
}
