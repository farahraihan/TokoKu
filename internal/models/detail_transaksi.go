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
	// Hapus semua detail transaksi dari database
	result := dm.db.Where("transaksi_id = ?", transaksiID).Delete(&DetailTransaksi{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete detail transaksi: %w", result.Error)
	}

	// Pengecekan rows affected
	if result.RowsAffected == 0 {
		return fmt.Errorf("no detail transaksi deleted for transaksi ID %d", transaksiID)
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
