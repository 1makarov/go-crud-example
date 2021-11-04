package v1

import (
	"github.com/1makarov/go-crud-example/internal/delivery/http/v1/auth"
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
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			a := v1.Group("/auth")
			{
				auth.InitRouter(h.services.Auth, a)
			}

			b := v1.Group("/books", h.validAuth)
			{
				books.InitRouter(h.services.Books, b)
			}
		}
	}

	return router
}
