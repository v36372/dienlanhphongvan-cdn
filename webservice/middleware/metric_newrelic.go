package middleware

import (
	"dienlanhphongvan-cdn/metric"
	"fmt"

	"github.com/gin-gonic/gin"
)

func MetricWithNewrelic(transactionName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if metric.NewRelic().Enable {
			txn := metric.NewRelic().StartTransaction(fmt.Sprintf("%s | %s", c.Request.Method, transactionName), c.Writer, c.Request)
			defer txn.End()
			//Before logic
			c.Next()

			if err := c.Errors.Last(); err != nil {
				txn.NoticeError(err)
			}
		} else {
			c.Next()
		}
	}
}
