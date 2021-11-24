package v1

import (
	"github.com/1makarov/go-crud-example/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InitUsersRouter(v1 *gin.RouterGroup) {
	users := v1.Group("/users")
	{
		users.POST("/sign-up", h.signUp)
		users.POST("/sign-in", h.signIn)
	}
}

// SignUp Sign Up
// @Summary SignUp
// @Tags auth
// @ID sign-up
// @Param input body types.SignUpInput true " "
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/users/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input types.SignUpInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.Users.SignUp(c, input); err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}

// SignIn Sign In
// @Summary SignIn
// @Tags auth
// @ID sign-in
// @Param input body types.SignInInput true " "
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/users/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input types.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	token, err := h.services.Users.SignIn(c, input)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
