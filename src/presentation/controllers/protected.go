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
	hi := fmt.Sprintf("Hello, %v 👋!", request.User["role"])
	return protocols.HTTPResponse{
		StatusCode: 200,
		Data: map[string]interface{}{
			"data": hi,
		},
	}
}
