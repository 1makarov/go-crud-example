package v1

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) identity(c *gin.Context) {
	session := sessions.Default(c)

	switch v := session.Get("auth").(type) {
	case bool:
		if v != true {
			newResponse(c, http.StatusInternalServerError, fmt.Errorf("error auth"))
			return
		}
	default:
		newResponse(c, http.StatusInternalServerError, fmt.Errorf("error auth"))
		return
	}
}
