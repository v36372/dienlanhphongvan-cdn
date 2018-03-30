package middleware

import (
	"dienlanhphongvan-cdn/errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	applicationJSON = "application/json; charset=utf-8"
)

func RewriteStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var (
			status int = c.Writer.Status()
			err    error
			msg    string
		)

		if e := c.Errors.Last(); e != nil && e.Err != nil {
			if value, ok := e.Err.(errors.ErrInfo); ok {
				err = value.GetError()
			} else {
				err = e.Err
			}
		}

		switch status {
		case http.StatusOK, http.StatusCreated, http.StatusNotModified, http.StatusPartialContent:
			return

		case http.StatusBadRequest:
			msg = "bad params"

		case http.StatusNotFound:
			msg = "not found"

		case http.StatusForbidden:
			msg = "forbidden"

		case http.StatusUnauthorized:
			msg = "unauthorized"

		case http.StatusInternalServerError:
			msg = "internal server error"

		default:
			msg = "unknown"
		}
		// reset
		c.Errors = nil
		// rewrite
		/*if err != nil {
			c.JSON(status, respWithErrors(msg, err))
		} else {*/
		c.JSON(status, respWithErrorMessage(msg))
		fmt.Println(err)
		//}
	}
}

type responseError struct {
	Message string `json:"message,omitempty"`
	Errors  error  `json:"errors,omitempty"`
}

func respWithErrorMessage(msg string) responseError {
	return responseError{
		Message: msg,
	}
}

func respWithErrors(msg string, err error) responseError {
	return responseError{
		Message: msg,
		Errors:  err,
	}
}
