package httpresponse

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func Default(msg string, status int) *Response {
	return &Response{msg, status}
}
