package client

import (
	"fmt"
	"net/http"
)

type ErrorsResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"`          // error message
	Errors   []Error        `json:"errors,omitempty"` // more detail on individual errors
}

func (r ErrorsResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Errors)
}
