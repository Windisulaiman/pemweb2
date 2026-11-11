package model

import "gorm.io/gorm"

type Matakuliah struct {
	gorm.Model
	Kode     string `gorm:"unique;size:10;not null" json:"kode"`
	Nama     string `gorm:"size:100;not null" json:"nama"`
	Sks      int    `gorm:"not null" json:"sks"`
	Semester int    `gorm:"not null" json:"semester"`
}
