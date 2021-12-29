package models

import "github.com/jinzhu/gorm"

type WaliKelass struct {
	gorm.Model
	Kelas                    int    `gorm:"not null" form:"kelas" json:"kelas"`
	Nama                     string `gorm:"not null" form:"nama" json:"nama"`
	TempatTanggalLahirTempat string `gorm:"not null" form:"tempat_tanggal_lahir_tempat" json:"tempat_tanggal_lahir_tempat"`
	TempatTanggalLahirTgl    string `gorm:"not null" form:"tempat_tanggal_lahir_tgl" json:"tempat_tanggal_lahir_tgl"`
	TempatTanggalLahirBln    string `gorm:"not null" form:"tempat_tanggal_lahir_bln" json:"tempat_tanggal_lahir_bln"`
	TempatTanggalLahirTahun  string `gorm:"not null" form:"tempat_tanggal_lahir_tahun" json:"tempat_tanggal_lahir_tahun"`
	Keluar                   string `gorm:"not null" form:"keluar" json:"keluar"`
}

func (WaliKelass) TableName() string {
	return "tbl_wali_kelas"
}
