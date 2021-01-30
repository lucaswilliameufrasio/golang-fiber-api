package controllers

import protocols "lucaswilliameufrasio/golang-fiber-api/src/presentation/protocols"

func NewStatusController() protocols.Controller {
	return StatusController{}
}

type StatusController struct{}

// Handler is a controller to handle request and respond with a great message
func (sts StatusController) Handler(request *protocols.HTTPRequest) protocols.HTTPResponse {
	return protocols.HTTPResponse{
		StatusCode: 200,
		Data: map[string]interface{}{
			"data": "Hello, World 👋!",
		},
	}
}
