package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Halo Pemrograman Web II")
	})

	// Langkah 7: Endpoint JSON dasar
	app.Get("/api/info", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"aplikasi": "Latihan Fiber",
			"versi":    "1.0.0",
			"status":   "berjalan",
		})
	})

	// Tugas 1: Endpoint data mahasiswa
	app.Get("/api/mahasiswa", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"nim":   "H1H024005", // Ubah sesuai NIM Anda
			"nama":  "Windi Sulaiman Ismansa",  // Ubah sesuai Nama Anda
			"prodi": "Teknik Komputer",
		})
	})

	log.Fatal(app.Listen(":3000"))
}