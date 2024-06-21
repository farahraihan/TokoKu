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

func (bm *BarangModel) GetBarang() ([]Barang, error) {
	var barangs []Barang
	err := bm.db.Find(&barangs).Error

	if err != nil {
		return nil, err
	}

	return barangs, nil
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

func (bm *BarangModel) GetBarangByID(barangID uint) (*Barang, error) {
	var barang Barang
	if err := bm.db.First(&barang, barangID).Error; err != nil {
		return nil, err
	}
	return &barang, nil
}

func (bm *BarangModel) UpdateBarang(id uint, updatedData Barang) (Barang, error) {
	var existingData Barang
	// Cari data pelanggan berdasarkan ID
	err := bm.db.First(&existingData, id).Error
	if err != nil {
		return Barang{}, err
	}

	// Update data yang ditemukan
	existingData.Nama = updatedData.Nama
	existingData.Stok = updatedData.Stok
	existingData.Harga = updatedData.Harga

	// Simpan data yang telah diperbarui kembali ke database
	err = bm.db.Save(&existingData).Error
	if err != nil {
		return Barang{}, err
	}

	return existingData, nil
}