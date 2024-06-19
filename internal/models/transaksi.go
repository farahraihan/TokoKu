package models

import (
	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model
	PegawaiId        uint
	CustomerId       uint
	DetailTransaksis []DetailTransaksi `gorm:"foreignKey:TransaksiID"`
}

type TransaksiModel struct {
	db *gorm.DB
}

func NewTransaksiModel(connection *gorm.DB) *TransaksiModel {
	return &TransaksiModel{
		db: connection,
	}
}
