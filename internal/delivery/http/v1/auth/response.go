package auth

import "github.com/gin-gonic/gin"

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, err error) {
	c.AbortWithStatusJSON(statusCode, response{err.Error()})
}
