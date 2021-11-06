package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, err error) {
	logrus.Error(err)
	c.AbortWithStatusJSON(statusCode, response{err.Error()})
}
