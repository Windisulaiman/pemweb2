package handler

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/repository"
)

func DaftarProgramStudi(c fiber.Ctx) error {
	daftar, err := repository.AmbilProgramStudiDenganJumlah()
	if err != nil {
		log.Println("[ERROR] AmbilProgramStudiDenganJumlah:", err) // ← tambahkan
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
	})
}
