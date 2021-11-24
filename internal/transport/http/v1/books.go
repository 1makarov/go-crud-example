package v1

import (
	"fmt"
	"github.com/1makarov/go-crud-example/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) InitBooksRouter(v1 *gin.RouterGroup) {
	auth := v1.Group("/books", h.identity)
	{
		auth.POST("/", h.Create)
		auth.GET("/:id", h.GetByID)
		auth.GET("/", h.GetAll)
		auth.DELETE("/:id", h.DeleteByID)
		auth.PUT("/:id", h.UpdateByID)
	}
}

// Create Book
// @Summary Create
// @Security AuthKey
// @Tags book
// @ID create-book
// @Param input body types.BookCreateInput true " "
// @Success 201 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/books/ [post]
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

// GetByID
// @Summary Get by id
// @Security AuthKey
// @Tags book
// @Description get book by id
// @ID get-book-by-id
// @Param id path string true " "
// @Success 200 {object} types.Book
// @Failure 400,404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Failure default {object} types.ErrorResponse
// @Router /api/v1/books/{id} [get]
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

// GetAll
// @Summary Get all
// @Security AuthKey
// @Tags book
// @ID get-all-books
// @Success 200 {object} []types.Book
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/books/ [get]
func (h *Handler) GetAll(c *gin.Context) {
	books, err := h.services.Books.GetAll(c.Request.Context())
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, books)
}

// DeleteByID
// @Summary Delete by id
// @Security AuthKey
// @Tags book
// @ID delete-book-by-id
// @Param id path string true " "
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/books/{id} [delete]
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

// UpdateByID
// @Summary Update by id
// @Security AuthKey
// @Tags book
// @ID update-book-by-id
// @Param id path string true " "
// @Param input body types.BookUpdateInput true " "
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/books/{id} [put]
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
