package handler

import (
	"sort" // Diimpor untuk melakukan sorting slice
	"strconv"

	"latihan-fiber-api/internal/model"

	"github.com/gofiber/fiber/v3"
)

func DaftarMahasiswa(c fiber.Ctx) error {
	angkatan := c.Query("angkatan", "")
	urut := c.Query("urut", "") // [TUGAS 2] Membaca parameter kueri "urut" (nama/ipk)

	hasil := []model.Mahasiswa{}

	// Filter berdasarkan angkatan jika parameter angkatan diisi[cite: 7]
	if angkatan == "" {
		hasil = append(hasil, model.Daftar...)
	} else {
		nilaiAngkatan, err := strconv.Atoi(angkatan)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Parameter angkatan harus berupa angka") //[cite: 7]
		}
		for i := 0; i < len(model.Daftar); i++ {
			if model.Daftar[i].Angkatan == nilaiAngkatan {
				hasil = append(hasil, model.Daftar[i])
			}
		}
	}

	//tugas nomor2
	if urut == "nama" {
		sort.Slice(hasil, func(i, j int) bool {
			return hasil[i].Nama < hasil[j].Nama
		})
	} else if urut == "ipk" {
		sort.Slice(hasil, func(i, j int) bool {
			return hasil[i].Ipk > hasil[j].Ipk
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   hasil,
	})
}

func DetailMahasiswa(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id")) //[cite: 7]
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka") //[cite: 7]
	}

	for i := 0; i < len(model.Daftar); i++ {
		if model.Daftar[i].ID == id {
			return c.JSON(fiber.Map{
				"sukses": true,
				"data":   model.Daftar[i], //[cite: 8]
			})
		}
	}

	return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan") //[cite: 8]
}

func DetailMahasiswaByNim(c fiber.Ctx) error {
	nim := c.Params("nim")

	for i := 0; i < len(model.Daftar); i++ {
		if model.Daftar[i].Nim == nim {
			return c.JSON(fiber.Map{
				"sukses": true,
				"data":   model.Daftar[i],
			})
		}
	}

	return fiber.NewError(fiber.StatusNotFound, "Mahasiswa dengan NIM tersebut tidak ditemukan")
}
