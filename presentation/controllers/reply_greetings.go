package controllers

import presentationprotocols "lucaswilliameufrasio/golang-fiber-api/presentation/protocols"

// ReplyGreetingsController is a controller to handle request and respond with a great message
func ReplyGreetingsController(request *presentationprotocols.HTTPRequest) presentationprotocols.HTTPResponse {
	return presentationprotocols.HTTPResponse{
		StatusCode: 200,
		Data: map[string]interface{}{
			"data": "Hello, World 👋!",
		},
	}
}
