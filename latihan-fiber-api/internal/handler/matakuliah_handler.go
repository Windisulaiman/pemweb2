package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

type PermintaanMatakuliah struct {
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	Sks      int    `json:"sks"`
	Semester int    `json:"semester"`
}

func DaftarMatakuliah(c fiber.Ctx) error {
	halaman, _ := strconv.Atoi(c.Query("halaman", "1"))
	perHalaman, _ := strconv.Atoi(c.Query("per_halaman", "10"))
	semester, _ := strconv.Atoi(c.Query("semester", "0"))

	if halaman < 1 {
		halaman = 1
	}
	if perHalaman < 1 || perHalaman > 100 {
		perHalaman = 10
	}

	filter := repository.FilterMatakuliah{
		Cari:       c.Query("cari", ""),
		Semester:   semester,
		Halaman:    halaman,
		PerHalaman: perHalaman,
	}

	daftar, total, err := repository.AmbilMatakuliah(filter)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data")
	}

	totalHalaman := int(total) / perHalaman
	if int(total)%perHalaman != 0 {
		totalHalaman++
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
		"meta": fiber.Map{
			"halaman":       halaman,
			"per_halaman":   perHalaman,
			"total":         total,
			"total_halaman": totalHalaman,
		},
	})
}

func DetailMatakuliah(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
	}

	mk, err := repository.AmbilMatakuliahPerId(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Matakuliah tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data")
	}

	return c.JSON(fiber.Map{"sukses": true, "data": mk})
}

func BuatMatakuliah(c fiber.Ctx) error {
	var p PermintaanMatakuliah
	if err := c.Bind().Body(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
	}

	mk := model.Matakuliah{Kode: p.Kode, Nama: p.Nama, Sks: p.Sks, Semester: p.Semester}
	if err := repository.SimpanMatakuliah(&mk); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Gagal menyimpan data, periksa keunikan kode")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data matakuliah berhasil dibuat",
		"data":   mk,
	})
}

func PerbaruiMatakuliah(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
	}

	mk, err := repository.AmbilMatakuliahPerId(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Matakuliah tidak ditemukan")
	}

	var p PermintaanMatakuliah
	if err := c.Bind().Body(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
	}

	if p.Kode != "" {
		mk.Kode = p.Kode
	}
	if p.Nama != "" {
		mk.Nama = p.Nama
	}
	if p.Sks != 0 {
		mk.Sks = p.Sks
	}
	if p.Semester != 0 {
		mk.Semester = p.Semester
	}

	if err := repository.PerbaruiMatakuliah(&mk); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Gagal memperbarui data")
	}

	return c.JSON(fiber.Map{"sukses": true, "pesan": "Data matakuliah berhasil diperbarui", "data": mk})
}

func HapusMatakuliah(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
	}

	mk, err := repository.AmbilMatakuliahPerId(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Matakuliah tidak ditemukan")
	}

	if err := repository.HapusMatakuliah(&mk); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menghapus data")
	}

	return c.JSON(fiber.Map{"sukses": true, "pesan": "Data matakuliah berhasil dihapus"})
}
