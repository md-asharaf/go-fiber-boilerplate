package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
)

// HealthCheck provides health check endpoint for Fiber
func HealthCheck(c *fiber.Ctx) error {
	return utils.WriteSuccessResponse(c, nil, "ok")
}
