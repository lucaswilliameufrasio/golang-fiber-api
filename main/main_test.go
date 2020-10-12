// +build test

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 2, 2, "Expecting 2")
}

func TestMainRoute(t *testing.T) {
	app := server()

	// http.Request
	req := httptest.NewRequest("GET", "http://localhost:7777", nil)

	// http.Response
	resp, _ := app.Test(req)

	// Do something with results:
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body)) // => Hello, World 👋!

		t.Errorf("got %d, want %d", resp.StatusCode, 200)
	}
}

func TestV2MainRoute(t *testing.T) {
	app := server()

	// http.Request
	req := httptest.NewRequest("GET", "http://localhost:7777/api/v2/", nil)

	// http.Response
	resp, _ := app.Test(req)

	// Do something with results:
	body, _ := ioutil.ReadAll(resp.Body)
	// if resp.StatusCode != 200 {
	// fmt.Println(string(body)) // => Hello, World 👋!

	// 	t.Errorf("got %d, want %d", resp.StatusCode, 200)
	// }

	assert.Equal(t, "Hello, World 👋!", string(body))
}

func TestCorrectNameAndAge(t *testing.T) {
	app := server()

	name := "doe"
	age := "32"

	req := httptest.NewRequest(http.MethodGet, "http://localhost:7777/api/v1/profile/"+name+"/"+age, nil)

	resp, _ := app.Test(req)

	expectedResponse := `{"active":true,"info":"👴 ` + name + ` is ` + age + ` years old"}`

	body, _ := ioutil.ReadAll(resp.Body)
	responseBody := string(body)

	assert.Equal(t, expectedResponse, responseBody)
}

func TestCorrectEmailAndPassword(t *testing.T) {
	app := server()

	type User struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
	}

	type Response struct {
		Token string `json:"token"`
		User  *User  `json:"user"`
	}

	// http.Request
	httpRequest := []byte(`{"email":"john@example.com","password":"doe"}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v2/login", bytes.NewBuffer(httpRequest))
	req.Header.Set("Content-Type", "application/json")

	// http.Response
	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	result := &Response{
		User: &User{},
	}
	err := json.Unmarshal(body, result)

	if err != nil {
		fmt.Println(err)
	}

	userId := result.User.Id
	userEmail := result.User.Email

	assert.Equal(t, 1, userId)
	assert.Equal(t, "john@example.com", userEmail)
}

func TestProtectedRoute(t *testing.T) {
	app := server()

	type User struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
	}

	type Response struct {
		Token string `json:"token"`
		User  *User  `json:"user"`
	}

	// http.Request
	httpRequest := []byte(`{"email":"john@example.com","password":"doe"}`)

	reqLogin := httptest.NewRequest(http.MethodPost, "http://localhost:7777/api/v2/login", bytes.NewBuffer(httpRequest))
	reqLogin.Header.Set("Content-Type", "application/json")

	respWithToken, _ := app.Test(reqLogin)

	body, _ := ioutil.ReadAll(respWithToken.Body)

	result := &Response{
		User: &User{},
	}
	err := json.Unmarshal(body, result)

	if err != nil {
		fmt.Println(err)
	}

	accessToken := result.Token

	mainRequest := httptest.NewRequest(http.MethodGet, "http://localhost:7777/api/v2/profile", nil)
	mainRequest.Header.Set("Content-Type", "application/json")
	mainRequest.Header.Set("Authorization", "Bearer "+accessToken)

	responseExpected, _ := app.Test(mainRequest)
	expectedBody, _ := ioutil.ReadAll(responseExpected.Body)

	assert.Equal(t, "Welcome, John Doe", string(expectedBody))
}
