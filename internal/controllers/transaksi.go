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

func (tc *TransaksiController) AddTransaksi(id uint) (bool, error) {
	var newData models.Transaksi
	var barangId uint
	var jumlah uint
	var harga float64
	var menu int

	fmt.Print("Masukkan id customer: ")
	fmt.Scanln(&newData.CustomerID)
	newData.PegawaiID = id

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah barang")
		fmt.Println("2. Selesai")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&menu)

		if menu == 2 {
			break // Keluar dari loop jika memilih menu 2 (selesai)
		}

		fmt.Print("Masukkan id barang: ")
		fmt.Scanln(&barangId)
		fmt.Print("Masukkan jumlah barang: ")
		fmt.Scanln(&jumlah)
		fmt.Print("Masukkan harga barang: ")
		fmt.Scanln(&harga)

		newTransactionID, err := tc.model.AddTransaksi(newData)
		if err != nil {
			return false, err // Mengembalikan false karena tidak ada transaksi yang berhasil ditambahkan
		}

		if tc.DetailTransaksiController != nil {
			err = tc.DetailTransaksiController.AddDetailTransaksi(newTransactionID, barangId, jumlah, harga)
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil // Mengembalikan true jika proses penambahan transaksi berhasil selesai
}
