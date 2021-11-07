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

// CreateToken
// @Summary Create
// @Tags auth
// @Description create auth token
// @ID create-auth-token
// @Produce json
// @Success 200 {string} string "token"
// @Failure 400,404 {object} response
// @Failure default {object} response
// @Router /api/v1/auth/create [get]
func (h *Handler) CreateToken(c *gin.Context) {
	token, err := h.services.Auth.CreateToken()
	if err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
