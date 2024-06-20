package models

import (
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
	err := pm.db.Create(&newData).Error
	if err != nil {
		return Pegawai{}, err
	}
	return newData, nil
}

// UpdatePegawaiByID updates a Pegawai record in the database by ID
func (m *PegawaiModel) UpdatePegawaiByID(id uint, newData Pegawai) error {
	var pegawai Pegawai
	if err := m.db.First(&pegawai, id).Error; err != nil {
		return err
	}

	// Update the fields that are allowed to be updated
	pegawai.Username = newData.Username
	pegawai.Nama = newData.Nama
	pegawai.Gender = newData.Gender
	pegawai.NoTelp = newData.NoTelp
	pegawai.Email = newData.Email
	pegawai.Password = newData.Password

	if err := m.db.Save(&pegawai).Error; err != nil {
		return err
	}

	return nil
}
