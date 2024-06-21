package controllers

import (
	"fmt"
	"strconv"
	"tokoku/internal/models"
)

type BarangController struct {
	model *models.BarangModel
}

func NewBarangController(m *models.BarangModel) *BarangController {
	return &BarangController{
		model: m,
	}
}

func (bc *BarangController) GetBarang() {
	barangs, err := bc.model.GetBarang()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, barang := range barangs {
		fmt.Println("----------------------------------")
		fmt.Printf("ID: %d\n", barang.ID)
		fmt.Printf("PegawaiID: %d\n", barang.PegawaiID)
		fmt.Printf("Nama: %s\n", barang.Nama)
		fmt.Printf("Stok: %d\n", barang.Stok)
		fmt.Printf("Harga: %.2f\n", barang.Harga)
		fmt.Println()

	}
}

func (bc *BarangController) DecreaseStock(barangID uint, quantity uint) error {
	err := bc.model.DecreaseStock(barangID, quantity)
	if err != nil {
		return fmt.Errorf("failed to decrease stock: %w", err)
	}
	return nil
}

func (bc *BarangController) IncreaseStock(barangID uint, quantity uint) error {
	err := bc.model.IncreaseStock(barangID, quantity)
	if err != nil {
		return fmt.Errorf("failed to increase stock: %w", err)
	}
	return nil
}

func (bc *BarangController) GetBarangByID(barangID uint) (*models.Barang, error) {
	barang, err := bc.model.GetBarangByID(barangID)
	if err != nil {
		return nil, fmt.Errorf("failed to get barang: %w", err)
	}
	return barang, nil
}

func (bc *BarangController) UpdateStock() error {
	var barangIDInput string
	var quantityInput string

	// Meminta input ID barang dari pengguna
	fmt.Print("Masukkan ID barang: ")
	_, err := fmt.Scanln(&barangIDInput)
	if err != nil {
		return fmt.Errorf("gagal membaca ID barang: %w", err)
	}

	// Meminta input jumlah stok yang ingin ditambahkan dari pengguna
	fmt.Print("Masukkan jumlah stok yang ingin ditambahkan: ")
	_, err = fmt.Scanln(&quantityInput)
	if err != nil {
		return fmt.Errorf("gagal membaca jumlah stok: %w", err)
	}

	// Konversi input menjadi tipe data uint
	barangID, err := strconv.ParseUint(barangIDInput, 10, 32)
	if err != nil {
		return fmt.Errorf("ID barang tidak valid: %w", err)
	}
	quantity, err := strconv.ParseUint(quantityInput, 10, 32)
	if err != nil {
		return fmt.Errorf("jumlah stok tidak valid: %w", err)
	}

	// Panggil metode IncreaseStock dari model
	err = bc.model.IncreaseStock(uint(barangID), uint(quantity))
	if err != nil {
		return fmt.Errorf("gagal menambah stok: %w", err)
	}

	fmt.Println("Stok berhasil ditambahkan.")
	return nil
}

func (bc *BarangController) DeleteBarang() (bool, error) {
	var deleteData models.Barang
	fmt.Print("Masukkan ID Barang yang ingin dihapus: ")
	fmt.Scanln(&deleteData.ID)
	err := bc.model.DeleteBarang(deleteData.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}