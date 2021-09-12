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

func SUTLogin() *fiber.App {
	return config.App(fiber.Config{
		DisableStartupMessage: true,
	})
}

func TestCorrectEmailAndPassword(t *testing.T) {
	app := SUTLogin()

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
	app := SUTLogin()

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
	app := SUTLogin()

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

func TestMissingEmail(t *testing.T) {
	app := SUTLogin()

	// http.Request
	httpRequest := []byte(`{"email":"","password":"pass"}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v1/login", bytes.NewBuffer(httpRequest))
	req.Header.Set("Content-Type", "application/json")

	// http.Response
	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	expectedResponse := `{"error":"Missing param: Email"}`

	assert.Equal(t, expectedResponse, string(body))
}

func TestMissingPassword(t *testing.T) {
	app := SUTLogin()

	// http.Request
	httpRequest := []byte(`{"email":"user1@example.com","password":""}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v1/login", bytes.NewBuffer(httpRequest))
	req.Header.Set("Content-Type", "application/json")

	// http.Response
	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	expectedResponse := `{"error":"Missing param: Password"}`

	assert.Equal(t, expectedResponse, string(body))
}

func TestNilBody(t *testing.T) {
	app := SUTLogin()

	req := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v1/login", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	expectedResponse := `{"error":"Missing param: Email"}`

	assert.Equal(t, expectedResponse, string(body))
}

func TestNilParameters(t *testing.T) {
	app := SUTLogin()

	httpRequest := []byte(`{}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v1/login", bytes.NewBuffer(httpRequest))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	expectedResponse := `{"error":"Missing param: Email"}`

	assert.Equal(t, expectedResponse, string(body))
}
