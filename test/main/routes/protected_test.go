package routes_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lucaswilliameufrasio/golang-fiber-api/src/main/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func SUTProtected() *fiber.App {
	return config.App(fiber.Config{
		DisableStartupMessage: true,
	})
}

func TestProtectedRoute(t *testing.T) {
	app := SUTProtected()

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
	app := SUTProtected()

	req := httptest.NewRequest(http.MethodGet, "http://localhost:7777/api/v1/protected", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	expectedResponse := `{"error":"Unauthorized"}`

	assert.Equal(t, expectedResponse, string(body))
}
