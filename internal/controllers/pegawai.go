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

// UpdatePegawai updates an existing pegawai record by ID
func (c *PegawaiController) UpdatePegawai() {
	var newData models.Pegawai

	fmt.Print("Masukkan ID Pegawai yang ingin diupdate: ")
	var id uint
	fmt.Scanln(&id)

	// Input new data
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&newData.Username)
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&newData.Nama)
	fmt.Print("Masukkan Gender (L/P): ")
	fmt.Scanln(&newData.Gender)
	fmt.Print("Masukkan Nomor Telepon: ")
	fmt.Scanln(&newData.NoTelp)
	fmt.Print("Masukkan Email: ")
	fmt.Scanln(&newData.Email)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&newData.Password)

	// Update pegawai by ID
	err := c.model.UpdatePegawaiByID(id, newData)
	if err != nil {
		fmt.Printf("Gagal melakukan update pegawai: %v\n", err)
		return
	}

	fmt.Println("Data pegawai berhasil diupdate!")
}
