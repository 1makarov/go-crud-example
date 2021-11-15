package v1

import (
	_ "github.com/1makarov/go-crud-example/docs"
	"github.com/1makarov/go-crud-example/internal/pkg/auth"
	"github.com/1makarov/go-crud-example/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *services.Service
	manager  *auth.Manager
}

func NewHandler(services *services.Service, manager *auth.Manager) *Handler {
	return &Handler{services: services, manager: manager}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			h.InitAuthRouter(v1)
			h.InitBooksRouter(v1)
		}
	}

	return router
}
