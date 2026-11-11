package repository

import "latihan-fiber-api/internal/config"

type ProgramStudiDenganJumlah struct {
	ID              uint   `json:"id"`
	Kode            string `json:"kode"`
	Nama            string `json:"nama"`
	Jenjang         string `json:"jenjang"`
	JumlahMahasiswa int64  `json:"jumlah_mahasiswa"`
}

func AmbilProgramStudiDenganJumlah() ([]ProgramStudiDenganJumlah, error) {
	var hasil []ProgramStudiDenganJumlah

	err := config.DB.
		Table("program_studis").
		Select(`
			program_studis.id,
			program_studis.kode,
			program_studis.nama,
			program_studis.jenjang,
			COUNT(mahasiswas.id) AS jumlah_mahasiswa
		`).
		Joins("LEFT JOIN mahasiswas ON mahasiswas.program_studi_id = program_studis.id AND mahasiswas.deleted_at IS NULL").
		Where("program_studis.deleted_at IS NULL").
		Group("program_studis.id, program_studis.kode, program_studis.nama, program_studis.jenjang").
		Order("program_studis.id ASC").
		Scan(&hasil).Error

	return hasil, err
}
