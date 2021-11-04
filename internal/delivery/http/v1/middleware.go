package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const errAuth = "error auth"

func (h *Handler) validAuth(c *gin.Context) {
	header := strings.Split(c.GetHeader("Authorization"), " ")
	if len(header) != 2 {
		newResponse(c, http.StatusUnauthorized, fmt.Errorf(errAuth))
		return
	}

	if err := h.services.Auth.ValidToken(header[1]); err != nil {
		newResponse(c, http.StatusUnauthorized, err)
		return
	}
}
