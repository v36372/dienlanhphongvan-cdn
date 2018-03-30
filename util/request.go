package util

import (
	"net/http"
	"net/http/httputil"
)

func DumpRequestHeader(req *http.Request) string {
	if req == nil {
		return ""
	}
	bytes, _ := httputil.DumpRequest(req, false)
	return string(bytes)
}
