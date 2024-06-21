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

func (pc *PegawaiController) AddPegawai() {
	var newData models.Pegawai

	fmt.Print("Masukkan username: ")
	fmt.Scanln(&newData.Username)
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&newData.Nama)
	fmt.Print("Masukkan gender (Pria/Wanita): ")
	var gender string
	fmt.Scanln(&gender)
	newData.Gender = models.Gender(gender)
	fmt.Print("Masukkan nomor telepon: ")
	fmt.Scanln(&newData.NoTelp)
	fmt.Print("Masukkan email: ")
	fmt.Scanln(&newData.Email)
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&newData.Password)

	// Panggil fungsi AddPegawai dari PegawaiModel untuk menyimpan data pegawai baru
	pegawai, err := pc.model.AddPegawai(newData)
	if err != nil {
		fmt.Printf("Gagal menambahkan pegawai: %v\n", err)
		return
	}

	fmt.Printf("Pegawai dengan ID %d berhasil ditambahkan.\n", pegawai.ID)
}
