package controllers

import (
	"fmt"
	"tokoku/internal/models"
)

type DetailTransaksiController struct {
	model            *models.DetailTransaksiModel
	barangController *BarangController
}

func NewDetailTransaksiController(m *models.DetailTransaksiModel, bc *BarangController) *DetailTransaksiController {
	return &DetailTransaksiController{
		model:            m,
		barangController: bc,
	}
}

func (dc *DetailTransaksiController) AddDetailTransaksi(transaksiID, barangID uint, jumlah uint, harga float64) error {
	// Kurangi stok barang
	err := dc.barangController.DecreaseStock(barangID, jumlah)
	if err != nil {
		return fmt.Errorf("failed to decrease stock: %w", err)
	}

	// Tambah detail transaksi ke database
	err = dc.model.AddDetailTransaksi(transaksiID, barangID, jumlah, harga)
	if err != nil {
		return fmt.Errorf("failed to add detail transaksi: %w", err)
	}
	return nil
}
