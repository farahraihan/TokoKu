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

func (pm *PegawaiModel) Login(username string, password string) (Pegawai, error) {
	var result Pegawai
	err := pm.db.Where("username = ? AND password = ?", username, password).First(&result).Error
	if err != nil {
		return Pegawai{}, err
	}
	return result, nil
}

func (pm *PegawaiModel) AddPegawai(newData Pegawai) (Pegawai, error) {
	err := pm.db.Create(&newData).Error
	if err != nil {
		return Pegawai{}, err
	}
	return newData, nil
}

func (pm *PegawaiModel) GetPegawai() ([]Pegawai, error) {
	var pegawais []Pegawai
	err := pm.db.Find(&pegawais).Error

	if err != nil {
		return nil, err
	}

	return pegawais, nil
}

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

func (pm *PegawaiModel) DeletePegawai(id uint) error {
	var pegawai Pegawai
	err := pm.db.First(&pegawai, id).Error
	if err != nil {
		return err
	}
	err = pm.db.Delete(&pegawai).Error
	if err != nil {
		return err
	}
	return nil
}
