package models

import "gorm.io/gorm"

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

func (tm *TransaksiModel) GetTransaksiWithDetails(transaksiID uint) (*Transaksi, error) {
	var transaksi Transaksi
	if err := tm.db.Preload("DetailTransaksis").First(&transaksi, transaksiID).Error; err != nil {
		return nil, err
	}
	return &transaksi, nil
}

func (tm *TransaksiModel) AddTransaksi(newData Transaksi) (uint, error) {
	err := tm.db.Create(&newData).Error
	if err != nil {
		return 0, err
	}
	return newData.ID, nil
}

func (tm *TransaksiModel) DeleteTransaksi(id uint) error {
	var Transaksi Transaksi
	err := tm.db.First(&Transaksi, id).Error
	if err != nil {
		return err
	}
	err = tm.db.Delete(&Transaksi).Error
	if err != nil {
		return err
	}
	return nil
}
