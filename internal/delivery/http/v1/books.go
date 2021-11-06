package v1

import (
	"fmt"
	"github.com/1makarov/go-crud-example/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) InitBooksRouter(v1 *gin.RouterGroup) {
	auth := v1.Group("/books", h.validAuth)
	{
		auth.POST("/create", h.Create)
		auth.GET("/get/:id", h.GetByID)
		auth.GET("/get-all", h.GetAll)
		auth.DELETE("/delete/:id", h.DeleteByID)
		auth.POST("/update/:id", h.UpdateByID)
	}
}

func (h *Handler) Create(c *gin.Context) {
	var v types.BookCreateInput

	if err := c.BindJSON(&v); err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.Books.Create(c.Request.Context(), v); err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) GetByID(c *gin.Context) {
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

	book, err := h.services.Books.GetByID(c.Request.Context(), id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) GetAll(c *gin.Context) {
	books, err := h.services.Books.GetAll(c.Request.Context())
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) DeleteByID(c *gin.Context) {
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

	if err = h.services.Books.DeleteByID(c.Request.Context(), id); err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) UpdateByID(c *gin.Context) {
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

	if err = h.services.Books.UpdateByID(c.Request.Context(), id, book); err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
