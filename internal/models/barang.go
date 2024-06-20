package models

import "gorm.io/gorm"

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

func (bm *BarangModel) GetBarangByID(barangID uint) (*Barang, error) {
	var barang Barang
	if err := bm.db.First(&barang, barangID).Error; err != nil {
		return nil, err
	}
	return &barang, nil
}
