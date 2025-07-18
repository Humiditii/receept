package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"statusCode"`
	Error      bool        `json:"error"`
}

func Send(c *gin.Context, statusCode int, message string, data interface{}, isError bool) {
	res := Response{
		Message:    message,
		Data:       data,
		StatusCode: statusCode,
		Error:      isError,
	}

	c.JSON(statusCode, res)
	c.Abort()
}
