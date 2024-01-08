package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thn-lee/01-task-management-api/pkg/models"
)

func Index() fiber.Handler {
	return func(c *fiber.Ctx) error {
		responseForm := models.ResponseForm{
			Success: true,
			Result: fiber.Map{
				"ip":            c.IP(),
				"remote":        c.IPs(),
				"client":        string(c.Context().UserAgent()),
				"secure":        c.Secure(),
				"conn_time":     c.Context().ConnTime(),
				"response_time": c.Context().Time(),
				"version":       c.Params("version", "0"),
			},
		}
		return c.JSON(responseForm)
	}
}
