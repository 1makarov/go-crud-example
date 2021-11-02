package v1

import (
	"github.com/1makarov/go-crud-example/internal/delivery/http/v1/books"
	"github.com/1makarov/go-crud-example/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() *gin.Engine {
	api := gin.Default()

	v1 := api.Group("/v1")
	{
		books.InitRouter(h.services.Books, v1)
	}

	return api
}
