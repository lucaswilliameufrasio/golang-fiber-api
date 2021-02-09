package presentationprotocols

type getParams func(string, ...string) string

// HTTPResponse fields
type HTTPResponse struct {
	StatusCode int
	Data       interface{}
}

// HTTPRequest fields
type HTTPRequest struct {
	Body   interface{}
	Params getParams
	UserID *int
	Token  string
}
