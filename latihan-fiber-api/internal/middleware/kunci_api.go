package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func KunciApi(apiKeyValid string) fiber.Handler {
	return func(c fiber.Ctx) error {
		apiKeyHeader := c.Get("X-Api-Key")

		// Jika header kosong atau tidak sesuai dengan .env
		if apiKeyHeader == "" || apiKeyHeader != apiKeyValid {
			return fiber.NewError(fiber.StatusUnauthorized, "Akses ditolak: API Key tidak valid")
		}

		return c.Next()
	}
}
