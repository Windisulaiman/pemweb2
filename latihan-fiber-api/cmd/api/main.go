package main

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/middleware"
	"latihan-fiber-api/internal/router"
)

func main() {
	konf := config.Muat()
	app := fiber.New(fiber.Config{
		AppName:      konf.NamaAplikasi,
		ErrorHandler: penanganGalat,
	})

	// 1. Middleware Global Bawaan
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recoverer.New())
	app.Use(middleware.WaktuProses())

	// 2. [TUGAS 3] Pasang middleware KunciApi SEBELUM router
	app.Use(middleware.KunciApi(konf.ApiKey))

	// 3. Registrasi Router
	router.Daftarkan(app)

	// 4. [Langkah 10] Middleware Wildcard 404 (Dipasang Paling Akhir)
	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Endpoint tidak tersedia",
		})
	})

	log.Fatal(app.Listen(":" + konf.Port))
}

func penanganGalat(c fiber.Ctx, err error) error {
	kode := fiber.StatusInternalServerError
	pesan := "Terjadi kesalahan pada server"

	var galatFiber *fiber.Error
	if errors.As(err, &galatFiber) {
		kode = galatFiber.Code
		pesan = galatFiber.Message
	}

	return c.Status(kode).JSON(fiber.Map{
		"sukses": false,
		"pesan":  pesan,
	})
}
