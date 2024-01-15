package middleware

import (
	"fmt"
	"rest-review/helpers"
	"rest-review/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")

		if authToken == "" {
			utils.ErrorMessage(c, &utils.ErrUnauthorized, fmt.Errorf("auth token empty"))
			return
		}

		if err := helpers.ValidateToken(authToken); err != nil {
			utils.ErrorMessage(c, &utils.ErrUnauthorized, fmt.Errorf("auth token unauthorized"))
			return
		}

		c.Next()
	}
}
