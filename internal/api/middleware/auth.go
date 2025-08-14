package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/yourusername/go-backend-boilerplate/internal/services"
	"go.uber.org/zap"
)

// contextKey is a custom type for context keys in middleware
type contextKey string

const userContextKey contextKey = "user"

// JWTAuth middleware validates JWT, fetches user from DB, and sets user in context for downstream handlers
func JWTAuth(jwtService *services.JWTService, userService *services.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				zap.L().Warn("Missing or invalid Authorization header")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Parse JWT token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				zap.L().Warn("Invalid JWT token", zap.Error(err))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Fetch user from DB using claims.UserID
			user, err := userService.GetUserByID(claims.UserID)
			if err != nil {
				zap.L().Warn("User not found for JWT claims", zap.Error(err))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Set user object in context for downstream handlers using custom key
			ctx := context.WithValue(r.Context(), userContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// CORS middleware for handling Cross-Origin Resource Sharing
func CORS() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
