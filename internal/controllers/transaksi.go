package controllers

import (
	"fmt"
	"strconv"
	"tokoku/internal/models"
)

type TransaksiController struct {
	transaksiModel *models.TransaksiModel
	barangModel    *models.BarangModel
}

func NewTransaksiController(tm *models.TransaksiModel, bm *models.BarangModel) *TransaksiController {
	return &TransaksiController{
		transaksiModel: tm,
		barangModel:    bm,
	}
}

func (tc *TransaksiController) PrintNotaTransaksi() error {
	var transaksiIDStr string

	fmt.Print("Masukkan ID Transaksi (angka positif): ")
	fmt.Scanln(&transaksiIDStr)

	// Mengonversi input string ke uint64
	transaksiID64, err := strconv.ParseUint(transaksiIDStr, 10, 64)
	if err != nil {
		return fmt.Errorf("gagal mengonversi ID Transaksi: %v", err)
	}

	// Mengonversi uint64 ke uint
	transaksiID := uint(transaksiID64)

	// Memanggil fungsi untuk mencetak nota transaksi berdasarkan ID transaksi
	err = tc.printNotaTransaksiByID(transaksiID)
	if err != nil {
		return fmt.Errorf("gagal mencetak nota transaksi: %v", err)
	}

	return nil
}

func (tc *TransaksiController) printNotaTransaksiByID(transaksiID uint) error {
	// Memuat informasi transaksi dengan detail transaksi dari database
	transaksi, err := tc.transaksiModel.GetTransaksiWithDetails(transaksiID)
	if err != nil {
		return fmt.Errorf("gagal memuat transaksi dengan detail: %v", err)
	}

	// Menampilkan informasi nota transaksi
	fmt.Printf("Nota Transaksi\n")
	fmt.Println("-------------------------------------------")
	fmt.Printf("ID Transaksi: %d\n", transaksi.ID)
	fmt.Printf("Tanggal: %s\n", transaksi.CreatedAt.Format("02-01-2006"))
	fmt.Printf("ID Pegawai: %d\n", transaksi.PegawaiID)
	fmt.Printf("ID Customer: %d\n", transaksi.CustomerID)
	fmt.Printf("Detail Transaksi:\n")
	totalHarga := 0.0 // Untuk menghitung total harga seluruh transaksi
	for _, detail := range transaksi.DetailTransaksis {
		// Memuat informasi barang untuk setiap detail transaksi
		barang, err := tc.barangModel.GetBarangByID(detail.BarangID)
		if err != nil {
			return fmt.Errorf("gagal memuat informasi barang untuk detail transaksi: %v", err)
		}

		// Menampilkan detail transaksi dan nama barang
		fmt.Printf("  - Barang: %s\n", barang.Nama)
		fmt.Printf("    Jumlah: %d\n", detail.Jumlah)
		fmt.Printf("    Harga: %.2f\n", detail.Harga)
		fmt.Printf("    Total: %.2f\n", float64(detail.Jumlah)*detail.Harga)
		totalHarga += float64(detail.Jumlah) * detail.Harga
	}

	// Menampilkan total harga seluruh transaksi
	fmt.Printf("Total Harga: %.2f\n", totalHarga)

	return nil
}
