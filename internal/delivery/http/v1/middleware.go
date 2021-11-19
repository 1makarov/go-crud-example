package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"

	errEmptyAuth         = "empty auth header"
	errInvalidAuthHeader = "invalid auth header"
	errTokenEmpty        = "token is empty"
)

func (h *Handler) identity(c *gin.Context) {
	if err := h.parseAuthToken(c); err != nil {
		newResponse(c, http.StatusUnauthorized, fmt.Errorf("auth: %s", err.Error()))
		return
	}
}

func (h *Handler) parseAuthToken(c *gin.Context) error {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return fmt.Errorf(errEmptyAuth)
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return fmt.Errorf(errInvalidAuthHeader)
	}

	if len(headerParts[1]) == 0 {
		return fmt.Errorf(errTokenEmpty)
	}

	return h.manager.Validate(headerParts[1])
}
