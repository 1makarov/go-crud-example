package books

import (
	"fmt"
	"github.com/1makarov/go-crud-example/internal/types"
	"github.com/1makarov/go-crud-example/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	service services.Books
}

func InitRouter(service services.Books, books *gin.RouterGroup) {
	h := handler{service: service}

	{
		books.POST("/create", h.Create)
		books.GET("/get/:id", h.GetByID)
		books.GET("/get-all", h.GetAll)
		books.DELETE("/delete/:id", h.DeleteByID)
		books.POST("/update/:id", h.UpdateByID)
	}
}

func (h *handler) Create(c *gin.Context) {
	var v types.BookCreateInput

	if err := c.BindJSON(&v); err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.service.Create(c.Request.Context(), v); err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *handler) GetByID(c *gin.Context) {
	v, ok := c.Params.Get("id")
	if !ok {
		newResponse(c, http.StatusBadRequest, fmt.Errorf("empty id"))
		return
	}

	id, err := strconv.Atoi(v)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	book, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *handler) GetAll(c *gin.Context) {
	books, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *handler) DeleteByID(c *gin.Context) {
	v, ok := c.Params.Get("id")
	if !ok {
		newResponse(c, http.StatusBadRequest, fmt.Errorf("empty isbn"))
		return
	}

	id, err := strconv.Atoi(v)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	if err = h.service.DeleteByID(c.Request.Context(), id); err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *handler) UpdateByID(c *gin.Context) {
	v, ok := c.Params.Get("id")
	if !ok {
		newResponse(c, http.StatusBadRequest, fmt.Errorf("empty isbn"))
		return
	}

	id, err := strconv.Atoi(v)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	var book types.BookUpdateInput

	if err = c.BindJSON(&book); err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	if err = h.service.UpdateByID(c.Request.Context(), id, book); err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
