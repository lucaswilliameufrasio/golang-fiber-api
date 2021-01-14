package presentationprotocols

type getParams func(string, ...string) string

// HTTPResponse fields
type HTTPResponse struct {
	StatusCode int
	Data       map[string]interface{}
}

// HTTPRequest fields
type HTTPRequest struct {
	Body   interface{}
	Params getParams
	User   map[string]string
}
