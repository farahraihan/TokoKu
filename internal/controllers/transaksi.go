package controllers

import "tokoku/internal/models"

type TransaksiController struct {
	model *models.TransaksiModel
}

func NewTransaksiController(m *models.TransaksiModel) *TransaksiController {
	return &TransaksiController{
		model: m,
	}
}
