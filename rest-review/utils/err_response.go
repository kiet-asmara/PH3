package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ErrorMessage(c *gin.Context, apiError *APIError, err error) *gin.Context {
	c.Abort()
	c.JSON(apiError.Code, gin.H{"error": APIError{
		Code:    apiError.Code,
		Message: apiError.Message,
	}})
	log.Println(err)
	return c
}
