package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func ParseAndValidateInput(c *fiber.Ctx, i interface{}) error {
	if err := c.BodyParser(i); err != nil {
		return WriteErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}
	if err := ValidateStruct(i); err != nil {
		return WriteErrorResponse(c, fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}
	return nil
}
