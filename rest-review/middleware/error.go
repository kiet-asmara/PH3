package middleware

import (
	"fmt"
	"rest-review/utils"

	"github.com/gin-gonic/gin"
)

// catch panics
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				utils.ErrorMessage(c, &utils.ErrInternalServer, fmt.Errorf("handling panic"))
			}
		}()

		c.Next()
	}
}
