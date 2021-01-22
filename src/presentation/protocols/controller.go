package presentationprotocols

// Controller is a interface that define what methods we have in a controller
type Controller func(*HTTPRequest) HTTPResponse
