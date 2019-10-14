package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthProcess() gin.HandlerFunc {
	return func(c *gin.Context) {
		//this function for abort the process
		//c.Abort()
		//return

		c.Next()
	}
}
