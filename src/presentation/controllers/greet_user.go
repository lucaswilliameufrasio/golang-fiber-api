package controllers

import (
	presentationprotocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

// GreetUserController is intended to return greetings message
func GreetUserController(request *presentationprotocols.HTTPRequest) presentationprotocols.HTTPResponse {
	userName := request.User["name"]

	return presentationprotocols.HTTPResponse{
		StatusCode: 200,
		Data: map[string]interface{}{
			"data": "Welcome, " + userName,
		},
	}
}
