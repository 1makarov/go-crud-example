package v1

import (
	"fmt"
	"github.com/1makarov/go-crud-example/internal/types"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InitUsersRouter(v1 *gin.RouterGroup) {
	users := v1.Group("/users")
	{
		auth := users.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
			auth.GET("/log-out", h.logOut)
		}
	}
}

// @Summary Sign up
// @Tags auth
// @ID sign-up
// @Param input body types.SignUpInput true " "
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/users/auth/sign-up [post]
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

// @Summary Sign in
// @Tags auth
// @ID sign-in
// @Param input body types.SignInInput true " "
// @Success 200 "OK"
// @Failure 401 {object} types.ErrorResponse
// @Router /api/v1/users/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input types.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusUnauthorized, err)
		return
	}

	_, err := h.services.Users.SignIn(c, input)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err)
		return
	}

	if err = updateSession(c, true); err != nil {
		newResponse(c, http.StatusUnauthorized, fmt.Errorf("error create session: %w", err))
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Log out
// @Tags auth
// @ID log-out
// @Success 200 "OK"
// @Failure 401 {object} types.ErrorResponse
// @Router /api/v1/users/auth/log-out [get]
func (h *Handler) logOut(c *gin.Context) {
	if err := updateSession(c, false); err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Errorf("error clear session: %w", err))
		return
	}

	c.Status(http.StatusOK)
}

func updateSession(c *gin.Context, b bool) error {
	session := sessions.Default(c)
	session.Set("auth", b)

	if err := session.Save(); err != nil {
		return err
	}

	return nil
}
