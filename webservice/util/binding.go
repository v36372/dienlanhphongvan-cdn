package util

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindJSON(c *gin.Context, ret interface{}) error {
	if err := binding.JSON.Bind(c.Request, ret); err != nil {
		return err
	}
	return nil
}

func BindForm(c *gin.Context, ret interface{}) error {
	if err := binding.Form.Bind(c.Request, ret); err != nil {
		return err
	}
	return nil
}
