package middleware

import (
	"github.com/gofiber/fiber/v2"
	s "github.com/md-asharaf/go-fiber-boilerplate/internal/services"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
)

// CORS middleware for Fiber
func CORS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	}
}

// JWTAuth middleware for Fiber
func JWTAuth(jwtService *s.JWTService, userService *s.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return utils.WriteErrorResponse(c, fiber.StatusUnauthorized, "Missing or invalid Authorization header")
		}
		tokenString := authHeader[7:]
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return utils.WriteErrorResponse(c, fiber.StatusUnauthorized, "Invalid JWT token")
		}
		user, err := userService.GetUserByID(claims.UserID)
		if err != nil {
			return utils.WriteErrorResponse(c, fiber.StatusUnauthorized, "User not found for JWT claims")
		}
		c.Locals("user", user)
		return c.Next()
	}
}
