package controllers

import (
	"fmt"
	"tokoku/internal/models"
)

type TransaksiController struct {
	model                      *models.TransaksiModel
	DetailTransaksiController *DetailTransaksiController
}

func NewTransaksiController(m *models.TransaksiModel, dc *DetailTransaksiController) *TransaksiController {
	return &TransaksiController{
		model:                      m,
		DetailTransaksiController: dc,
	}
}

func (tc *TransaksiController) AddTransaksi(id uint) (uint, error) {
	var newData models.Transaksi
	
	fmt.Print("Masukkan id customer: ")
	fmt.Scanln(&newData.CustomerID)
	newData.PegawaiID = id

	newTransactionID, err := tc.model.AddTransaksi(newData)
	if err != nil {
		return 0, err
	}
	return newTransactionID, nil // Mengembalikan true jika proses penambahan transaksi berhasil selesai
}
