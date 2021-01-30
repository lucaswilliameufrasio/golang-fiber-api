package controllers

import protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"

// StatusController is a controller to handle request and respond with a great message
func StatusController(request *protocols.HTTPRequest) protocols.HTTPResponse {
	return protocols.HTTPResponse{
		StatusCode: 200,
		Data: map[string]interface{}{
			"data": "Hello, World 👋!",
		},
	}
}
