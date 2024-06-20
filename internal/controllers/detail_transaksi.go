package controllers

import (
	"tokoku/internal/models"
)

type DetailTransaksiController struct {
	model *models.DetailTransaksiModel
}


func NewDetailTransaksiController(m *models.DetailTransaksiModel) *DetailTransaksiController {
	return &DetailTransaksiController{
		model: m,
	}
}

