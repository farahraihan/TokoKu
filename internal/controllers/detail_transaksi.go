package controllers

import (
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
