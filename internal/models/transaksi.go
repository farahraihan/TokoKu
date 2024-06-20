package models

import (
	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model
	PegawaiID        uint
	CustomerID       uint
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

func (tm *TransaksiModel) AddTransaksi(newData Transaksi) (uint, error) {
	err := tm.db.Create(&newData).Error
	if err != nil {
		return 0, err
	}
	return newData.ID, nil
}
