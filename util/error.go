package util

import (
	"dienlanhphongvan-cdn/errors"

	"github.com/gin-gonic/gin"
)

func statusCode(err error) int {
	if info, ok := err.(errors.ErrInfo); ok {
		return info.GetCode()
	}
	return 500
}

func CaseError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	if err != nil {
		c.Status(statusCode(err))
		c.Error(err)
		c.Abort()
		return true
	}
	return false
}
