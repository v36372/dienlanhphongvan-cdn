package middleware

import (
	"bytes"
	"dienlanhphongvan-cdn/util"
	"time"
	"utilities/ulog"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func InfoLogger(log *ulog.Ulogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var start = time.Now()
		// call next
		c.Next()
		var (
			status  = c.Writer.Status()
			method  = c.Request.Method
			header  = util.DumpRequestHeader(c.Request)
			latency = time.Now().Sub(start)
			api     = c.Request.URL.RequestURI()
		)

		log.LogInfo(api, ulog.Fields{
			"status":  status,
			"method":  method,
			"header":  header,
			"latency": latency.String(),
		})
	}
}

func InfoLoggerWithResponseBody(log *ulog.Ulogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			start = time.Now()
			w     = &ResponseWriter{
				body:           bytes.NewBufferString(""),
				ResponseWriter: c.Writer,
			}
		)
		// call next
		c.Writer = w
		c.Next()
		var (
			status       = c.Writer.Status()
			method       = c.Request.Method
			header       = util.DumpRequestHeader(c.Request)
			latency      = time.Now().Sub(start)
			api          = c.Request.URL.RequestURI()
			responseBody = ""
		)
		if method == "POST" {
			responseBody = w.body.String()
		}

		log.LogInfo(api, ulog.Fields{
			"status":        status,
			"method":        method,
			"header":        header,
			"latency":       latency.String(),
			"response_body": responseBody,
		})
	}
}
