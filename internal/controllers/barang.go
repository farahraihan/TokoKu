package controllers

import (
	"fmt"
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
