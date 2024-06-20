package controllers

import (
	"fmt"
	"tokoku/internal/models"
)

type TransaksiController struct {
	model *models.TransaksiModel
}

func NewTransaksiController(m *models.TransaksiModel) *TransaksiController {
	return &TransaksiController{
		model: m,
	}
}
func (cc *TransaksiController) AddTransaksi(id uint) (uint, error) {
	var newData models.Transaksi
	
	fmt.Print("Masukkan id customer: ")
	fmt.Scanln(&newData.CustomerID)
	newData.PegawaiID = id
	// looping 

	// fmt.Print("Masukkan jumlah barang: ")
	// fmt.Scanln(&newData.Jumlah)
	// fmt.Print("Masukkan jumlah barang: ")
	// fmt.Scanln(&newData.Jumlah)

	newTransactionID, err := cc.model.AddTransaksi(newData)
	if err != nil {
		return 0, err
	}

	return newTransactionID, nil
}


