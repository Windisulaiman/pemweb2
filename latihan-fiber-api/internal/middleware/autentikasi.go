package middleware

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func Autentikasi() fiber.Handler {
	konf := config.Muat()
	return func(c fiber.Ctx) error {
		header := c.Get("Authorization")
		if header == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Header Authorization tidak ditemukan")
		}

		bagian := strings.Split(header, " ")
		if len(bagian) != 2 || bagian[0] != "Bearer" {
			return fiber.NewError(fiber.StatusUnauthorized, "Format header Authorization tidak sesuai")
		}

		tokenString := bagian[1]

		// Cek apakah token sudah masuk daftar hitam (logout)
		var blacklist model.TokenBlacklist
		errCek := config.DB.Where("token = ?", tokenString).First(&blacklist).Error
		if errCek == nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Token sudah dicabut, silakan login kembali")
		}
		if !errors.Is(errCek, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusInternalServerError, "Gagal memeriksa status token")
		}

		klaim, err := auth.PeriksaToken(tokenString, konf.JwtSecret)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Token tidak valid atau sudah kedaluwarsa")
		}

		id, err := strconv.ParseUint(klaim.Subject, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Isi token tidak sesuai")
		}

		c.Locals("id_pengguna", uint(id))
		c.Locals("peran", klaim.Peran)
		c.Locals("token", tokenString)
		if klaim.ExpiresAt != nil {
			c.Locals("exp", klaim.ExpiresAt.Time)
		}

		return c.Next()
	}
}

func HanyaPeran(peranDiizinkan ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		peran, cocok := c.Locals("peran").(string)
		if !cocok {
			return fiber.NewError(fiber.StatusForbidden, "Peran pengguna tidak diketahui")
		}

		for i := 0; i < len(peranDiizinkan); i++ {
			if peranDiizinkan[i] == peran {
				return c.Next()
			}
		}

		return fiber.NewError(fiber.StatusForbidden, "Anda tidak berwenang mengakses sumber daya ini")
	}
}
