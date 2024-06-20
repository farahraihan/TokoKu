package controllers

import (
	"fmt"
	"tokoku/internal/models"
)

type PegawaiController struct {
	model *models.PegawaiModel
}

func NewPegawaiController(m *models.PegawaiModel) *PegawaiController {
	return &PegawaiController{
		model: m,
	}
}

func (pc *PegawaiController) GetPegawai() {
	pegawais, err := pc.model.GetPegawai()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, pegawai := range pegawais {
		fmt.Println("----------------------------------")
		fmt.Printf("ID: %d\n", pegawai.ID)
		fmt.Printf("Username: %s\n", pegawai.Username)
		fmt.Printf("Nama: %s\n", pegawai.Nama)
		fmt.Printf("Email: %s\n", pegawai.Email)
		fmt.Printf("Gender: %s\n", pegawai.Gender)
		fmt.Printf("NoTelp: %s\n", pegawai.NoTelp)
		fmt.Println()

	}
}
