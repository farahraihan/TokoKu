package controllers

import (
	"errors"
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

func (pc *PegawaiController) Login() (models.Pegawai, error) {
	var username, password string
	fmt.Print("Masukkan username : ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password : ")
	fmt.Scanln(&password)
	result, err := pc.model.Login(username, password)
	if err != nil {
		return models.Pegawai{}, errors.New("terjadi masalah ketika login")
	}
	return result, nil
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

func (pc *PegawaiController) AddPegawai() {
}

func (c *PegawaiController) UpdatePegawai() {
}

func (pc *PegawaiController) DeletePegawai() (bool, error) {
	var deleteData models.Pegawai
	fmt.Print("Masukkan ID Pegawai yang ingin dihapus: ")
	fmt.Scanln(&deleteData.ID)
	err := pc.model.DeletePegawai(deleteData.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
