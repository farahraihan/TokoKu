package models

import (
	"gorm.io/gorm"
)

type DetailTransaksi struct {
	gorm.Model
	TransaksiID uint
	BarangID    uint
	Jumlah      uint
	Harga       float64
}

type DetailTransaksiModel struct {
	db *gorm.DB
}

func NewDetailTransaksiModel(connection *gorm.DB) *DetailTransaksiModel {
	return &DetailTransaksiModel{
		db: connection,
	}
}
