package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ParseAndValidateInput(c *fiber.Ctx, i interface{}) error {
	if err := c.BodyParser(i); err != nil {
		return WriteErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}
	if err := validate.Struct(i); err != nil {
		return WriteErrorResponse(c, fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}
	return nil
}
