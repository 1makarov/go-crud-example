package auth

import (
	"github.com/1makarov/go-crud-example/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service services.Auth
}

func InitRouter(service services.Auth, auth *gin.RouterGroup) {
	h := &handler{service: service}

	auth.GET("/create", h.Create)
}

func (h *handler) Create(c *gin.Context) {
	token, err := h.service.CreateToken()
	if err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}
