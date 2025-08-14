package handlers

import (
	"net/http"

	"github.com/yourusername/go-backend-boilerplate/internal/models"
	"github.com/yourusername/go-backend-boilerplate/internal/services"
	"github.com/yourusername/go-backend-boilerplate/internal/utils"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input models.RegisterInput
	if err := utils.BindAndValidate(r, &input); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Validation failed: "+err.Error())
		return
	}

	resp, err := h.authService.Register(input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusCreated, resp)
}

// Login handles user login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input models.LoginInput
	if err := utils.BindAndValidate(r, &input); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Validation failed: "+err.Error())
		return
	}

	resp, err := h.authService.Login(input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, resp)
}
