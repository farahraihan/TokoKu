package models

import (
	"fmt"

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
	db          *gorm.DB
	barangModel *BarangModel
}

func NewDetailTransaksiModel(connection *gorm.DB, barangModel *BarangModel) *DetailTransaksiModel {
	return &DetailTransaksiModel{
		db:          connection,
		barangModel: barangModel,
	}
}

func (dm *DetailTransaksiModel) DeleteDetailTransaksi(transaksiID uint) error {
	var detailTransaksis []DetailTransaksi
	if err := dm.db.Where("transaksi_id = ?", transaksiID).Find(&detailTransaksis).Error; err != nil {
		return fmt.Errorf("failed to find detail transaksi: %w", err)
	}

	for _, dt := range detailTransaksis {
		// Kembalikan stok barang
		err := dm.barangModel.IncreaseStock(dt.BarangID, dt.Jumlah)
		if err != nil {
			return fmt.Errorf("failed to increase barang stock: %w", err)
		}
	}

	// Hapus detail transaksi dari database
	if err := dm.db.Where("transaksi_id = ?", transaksiID).Delete(&DetailTransaksi{}).Error; err != nil {
		return fmt.Errorf("failed to delete detail transaksi: %w", err)
	}

	return nil
}

func (dm *DetailTransaksiModel) AddDetailTransaksi(transaksiID, barangID uint, jumlah uint, harga float64) error {
	// Cek dan kurangi stok barang
	err := dm.barangModel.DecreaseStock(barangID, jumlah)
	if err != nil {
		return err
	}

	// Tambah detail transaksi ke database
	detail := DetailTransaksi{
		TransaksiID: transaksiID,
		BarangID:    barangID,
		Jumlah:      jumlah,
		Harga:       harga,
	}

	if err := dm.db.Create(&detail).Error; err != nil {
		return err
	}
	return nil
}
