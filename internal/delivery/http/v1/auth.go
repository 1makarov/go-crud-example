package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InitAuthRouter(v1 *gin.RouterGroup) {
	auth := v1.Group("/auth")
	{
		auth.GET("/create", h.CreateToken)
	}
}

func (h *Handler) CreateToken(c *gin.Context) {
	token, err := h.services.Auth.CreateToken()
	if err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}
