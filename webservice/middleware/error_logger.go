package middleware

import (
	"dienlanhphongvan-cdn/errors"
	"dienlanhphongvan-cdn/util"
	"time"
	"utilities/ulog"

	"github.com/gin-gonic/gin"
)

func ErrorLogger(log *ulog.Ulogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var start = time.Now()
		c.Next()
		if e := c.Errors.Last(); e != nil && e.Err != nil {
			var (
				err        error
				stackTrace string
				status     = c.Writer.Status()
				method     = c.Request.Method
				header     = util.DumpRequestHeader(c.Request)
				latency    = time.Now().Sub(start)
				api        = c.Request.URL.RequestURI()
			)
			if value, ok := e.Err.(errors.ErrInfo); ok {
				err = value.GetError()
				stackTrace = value.GetStackTrace()
			} else {
				err = e.Err
			}
			log.LogError(api, ulog.Fields{
				"status":      status,
				"method":      method,
				"header":      header,
				"latency":     latency.String(),
				"err":         err.Error(),
				"stack_trace": stackTrace,
			})
		}
	}
}
