package handlers

import (
	"net/http"

	"github.com/yourusername/go-backend-boilerplate/internal/models"
	"github.com/yourusername/go-backend-boilerplate/internal/services"
	"github.com/yourusername/go-backend-boilerplate/internal/utils"
)

// UserHandler handles user-related requests
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Me returns the current authenticated user's info
func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	userValue := r.Context().Value("user")
	if userValue == nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "Unauthorized: missing user context")
		return
	}

	user, ok := userValue.(*models.User)
	if !ok {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "Unauthorized: invalid user context")
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, user.ToResponse())
}
