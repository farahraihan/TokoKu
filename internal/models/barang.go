package models

import (
	"errors"

	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	Nama             string
	PegawaiID        uint
	Stok             uint              `gorm:"type:int;not null"`
	Harga            float64           `gorm:"type:float;not null"`
	DetailTransaksis []DetailTransaksi `gorm:"foreignKey:BarangID"`
}

type BarangModel struct {
	db *gorm.DB
}

func NewBarangModel(connection *gorm.DB) *BarangModel {
	return &BarangModel{
		db: connection,
	}
}

func (bm *BarangModel) DecreaseStock(barangID uint, quantity uint) error {
	var barang Barang
	if err := bm.db.First(&barang, barangID).Error; err != nil {
		return err
	}

	if barang.Stok < quantity {
		return errors.New("not enough stock")
	}

	barang.Stok -= quantity
	return bm.db.Save(&barang).Error
}

func (bm *BarangModel) IncreaseStock(barangID uint, quantity uint) error {
	var barang Barang
	if err := bm.db.First(&barang, barangID).Error; err != nil {
		return err
	}

	barang.Stok += quantity
	return bm.db.Save(&barang).Error
}
