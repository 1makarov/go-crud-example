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
	token, err := h.parseAuthToken(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, fmt.Errorf("identity: %s", err.Error()))
		return
	}

	if err = h.services.Users.ParseToken(token); err != nil {
		newResponse(c, http.StatusUnauthorized, fmt.Errorf("identity: %s", err.Error()))
		return
	}
}

func (h *Handler) parseAuthToken(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", fmt.Errorf(errEmptyAuth)
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", fmt.Errorf(errInvalidAuthHeader)
	}

	if len(headerParts[1]) == 0 {
		return "", fmt.Errorf(errTokenEmpty)
	}

	return headerParts[1], nil
}
