package models

import (
	"errors"

	"gorm.io/gorm"
)

// Struktur model untuk pegawai
type Pegawai struct {
	gorm.Model
	Username   string
	Nama       string
	Gender     Gender
	NoTelp     string
	Email      string
	Password   string
	Customers  []Customer  `gorm:"foreignKey:PegawaiID"`
	Transaksis []Transaksi `gorm:"foreignKey:PegawaiID"`
	Barangs    []Barang    `gorm:"foreignKey:PegawaiID"`
}

type PegawaiModel struct {
	db *gorm.DB
}

func NewPegawaiModel(connection *gorm.DB) *PegawaiModel {
	return &PegawaiModel{
		db: connection,
	}
}

func (pm *PegawaiModel) AddPegawai(newData Pegawai) (Pegawai, error) {
	// Validasi data
	if newData.Username == "" || newData.Nama == "" || newData.Email == "" || newData.Password == "" {
		return Pegawai{}, errors.New("semua field wajib diisi")
	}

	// Tambahkan pegawai baru
	result := pm.db.Create(&newData)
	if result.Error != nil {
		return Pegawai{}, result.Error
	}

	// Cek jumlah baris yang terpengaruh
	if result.RowsAffected == 0 {
		return Pegawai{}, errors.New("gagal menambahkan pegawai")
	}

	return newData, nil
}
