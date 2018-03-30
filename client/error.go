package client

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Desc string `json:"desc"` // resource on which the error occurred
}

func (e *Error) Error() string {
	if e == nil {
		return "<nil>"
	}
	bytes, _ := json.Marshal(e)
	return string(bytes)
}

func ErrorNotFound(err error) bool {
	if e, ok := err.(*ErrorResponse); ok {
		return e.Response.StatusCode == http.StatusNotFound
	}
	if e, ok := err.(*ErrorsResponse); ok {
		return e.Response.StatusCode == http.StatusNotFound
	}
	return false
}
