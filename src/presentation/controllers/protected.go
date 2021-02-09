package controllers

import (
	"fmt"
	protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"
)

func NewProtectedController() protocols.Controller {
	return ProtectedController{}
}

type ProtectedController struct{}

// Handler is a controller to handle request and respond with a great message
func (sts ProtectedController) Handler(request *protocols.HTTPRequest) protocols.HTTPResponse {
	return protocols.HTTPResponse{
		StatusCode: 200,
		Data: map[string]interface{}{
			"data": fmt.Sprintf("Hello, Dude 👋! Your ID is %v", *request.UserID),
		},
	}
}
