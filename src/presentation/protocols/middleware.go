package presentationprotocols

// Middleware is a interface that define what methods we have in a middleware
type Middleware interface {
	Handler(*HTTPRequest) HTTPResponse
}
