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
	// 1. Muat konfigurasi dari .env
	konf := config.Muat()

	// 2. Verifikasi konfigurasi (sementara, untuk debug)
	log.Printf("Aplikasi: %s | Port: %s | Lingkungan: %s", konf.NamaAplikasi, konf.Port, konf.Lingkungan)
	log.Printf("DB: user=%q host=%s port=%s nama=%s", konf.DbUser, konf.DbHost, konf.DbPort, konf.DbNama)

	// 3. Koneksi database + migrasi + data awal
	config.HubungkanBasisData(konf)
	config.Migrasi()
	config.DataAwal()

	if config.DB == nil {
		log.Fatal("FATAL: config.DB masih nil setelah HubungkanBasisData. Periksa konfigurasi .env dan koneksi MySQL.")
	}

	// 4. Inisialisasi Fiber
	app := fiber.New(fiber.Config{
		AppName:      konf.NamaAplikasi,
		ErrorHandler: penanganGalat,
	})

	// 5. Middleware Global Bawaan
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recoverer.New())
	//app.Use(middleware.WaktuProses())

	// 6. [TUGAS 3] Pasang middleware KunciApi SEBELUM router
	app.Use(middleware.KunciApi(konf.ApiKey))

	// 7. Registrasi Router
	router.Daftarkan(app)

	// 8. Middleware Wildcard 404 (Dipasang Paling Akhir)
	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Endpoint tidak tersedia",
		})
	})

	// 9. Jalankan server
	log.Printf("Server berjalan di http://localhost:%s", konf.Port)
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

	if kode == fiber.StatusInternalServerError {
		log.Println("Galat internal:", err.Error())
	}

	return c.Status(kode).JSON(fiber.Map{
		"sukses": false,
		"pesan":  pesan,
	})
}
