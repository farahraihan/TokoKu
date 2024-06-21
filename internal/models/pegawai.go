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

func (pm *PegawaiModel) Login(username string, password string) (Pegawai, error) {
	var result Pegawai
	err := pm.db.Where("username = ? AND password = ?", username, password).First(&result).Error
	if err != nil {
		return Pegawai{}, err
	}
	return result, nil
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

func (pm *PegawaiModel) GetPegawai() ([]Pegawai, error) {
	var pegawais []Pegawai
	err := pm.db.Find(&pegawais).Error

	if err != nil {
		return nil, err
	}

	return pegawais, nil
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

func (pm *PegawaiModel) UpdatePegawai(id uint, updatedData Pegawai) (Pegawai, error) {
	// Cari pegawai berdasarkan ID
	var pegawai Pegawai
	err := pm.db.First(&pegawai, id).Error
	if err != nil {
		return Pegawai{}, err
	}

	// Update data pegawai
	pegawai.Username = updatedData.Username
	pegawai.Nama = updatedData.Nama
	pegawai.Gender = updatedData.Gender
	pegawai.NoTelp = updatedData.NoTelp
	pegawai.Email = updatedData.Email
	pegawai.Password = updatedData.Password

	// Simpan perubahan ke dalam database
	err = pm.db.Save(&pegawai).Error
	if err != nil {
		return Pegawai{}, err
	}

	return pegawai, nil
}
