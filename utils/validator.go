package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedFields string
	Tag          string
	Value        string
}

var validate = validator.New()

const BODY_DATA = "body_data"

// validateStruct
// Validate an payload to a given struct, return error if something wrong, and return empty if all passed.
func validateStruct(data any) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedFields = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

// InterceptRequest
// Check and validate the payload, will intercept if the validation fails.
// If success, a locals will be set with given payload which can decrease unnecessary use of another
// BodyParser. Read more info about locals: https://docs.gofiber.io/api/ctx#locals
func InterceptRequest(data any) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		errors := validateStruct(data)

		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Given payload is invalid!",
				"errors":  errors,
			})
		}

		c.Locals(BODY_DATA, data)

		return c.Next()
	}
}

// GetPayload
// Smartly get a payload from locals and map them into given struct.
func GetPayload[T comparable](c *fiber.Ctx) T {
	payload := *c.Locals(BODY_DATA).(*T)

	return payload
}
