package infrastructure

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/thn-lee/01-task-management-api/pkg/models"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
)

var customErrorHandler = func(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := http.StatusInternalServerError
	var source string

	if e, ok := err.(*models.Error); ok {
		// Override status code if helpers.Error type
		code = e.Code
		source, _ = e.Source.(string)
	}

	responseData := models.ResponseForm{
		Success: false,
		Errors:  []models.ResponseError{models.ResponseError(*utils.NewError(code, source, err.Error()))},
	}

	// Return statuscode with error message
	err = c.Status(code).JSON(responseData)
	if err != nil {
		// In case the JSON fails
		log.Printf("customErrorHandler: %+v", err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}
