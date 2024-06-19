package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username   string
	Nama       string
	Gender     string
	NoTelp     string
	Password   string
	Email      string
	Pegawais   []Pegawai   `gorm:"foreignKey:AdminID"`
	Customers  []Customer  `gorm:"foreignKey:PegawaiID"`
	Transaksis []Transaksi `gorm:"foreignKey:PegawaiID"`
	Barangs    []Barang    `gorm:"foreignKey:PegawaiID"`
}

type AdminModel struct {
	db *gorm.DB
}

func NewAdminModel(connection *gorm.DB) *AdminModel {
	return &AdminModel{
		db: connection,
	}
}
