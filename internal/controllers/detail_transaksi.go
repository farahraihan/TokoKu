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

func (dc *DetailTransaksiController) AddDetailTransaksi(transaksiID uint) error {
	for {
		var barangID uint
		var jumlah uint

		fmt.Print("Masukkan ID Barang (0 untuk selesai): ")
		fmt.Scanln(&barangID)

		if barangID == 0 {
			break
		}

		fmt.Print("Masukkan jumlah barang : ")
		fmt.Scanln(&jumlah)

		// Dapatkan harga persatuan barang dari controller BarangController
		barang, err := dc.barangController.GetBarangByID(barangID)
		if err != nil {
			return fmt.Errorf("failed to get barang: %w", err)
		}

		// Hitung harga dari harga persatuan barang dikali jumlah
		harga := barang.Harga * float64(jumlah)

		// Tambah detail transaksi ke database
		err = dc.model.AddDetailTransaksi(transaksiID, barangID, jumlah, harga)
		if err != nil {
			return fmt.Errorf("failed to add detail transaksi: %w", err)
		}

		fmt.Printf("Detail transaksi untuk barang ID %d berhasil ditambahkan.\n", barangID)

	}
	return nil
}