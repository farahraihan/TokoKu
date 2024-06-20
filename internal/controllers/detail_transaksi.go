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

func (dc *DetailTransaksiController) DeleteDetailTransaksi(transaksiID uint) error {
	err := dc.model.DeleteDetailTransaksi(transaksiID)
	if err != nil {
		return fmt.Errorf("failed to delete detail transaksi: %w", err)
	}
	return nil
}
