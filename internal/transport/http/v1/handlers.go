package v1

import (
	_ "github.com/1makarov/go-crud-example/docs"
	"github.com/1makarov/go-crud-example/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.InitBooksRouter(v1)
		h.InitUsersRouter(v1)
	}
}
