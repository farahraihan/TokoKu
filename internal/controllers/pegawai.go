package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func (pc *PegawaiController) AddPegawai() error {
	var newData models.Pegawai
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan username: ")
	newData.Username, _ = reader.ReadString('\n')
	newData.Username = strings.TrimSpace(newData.Username)

	fmt.Print("Masukkan nama: ")
	newData.Nama, _ = reader.ReadString('\n')
	newData.Nama = strings.TrimSpace(newData.Nama)

	fmt.Print("Masukkan gender (Pria/Wanita): ")
	gender, _ := reader.ReadString('\n')
	newData.Gender = models.Gender(strings.TrimSpace(gender))

	if newData.Gender != models.Pria && newData.Gender != models.Wanita {
		return fmt.Errorf("gender tidak valid. Harus 'Pria' atau 'Wanita'")
	}

	fmt.Print("Masukkan nomor telepon: ")
	newData.NoTelp, _ = reader.ReadString('\n')
	newData.NoTelp = strings.TrimSpace(newData.NoTelp)

	fmt.Print("Masukkan email: ")
	newData.Email, _ = reader.ReadString('\n')
	newData.Email = strings.TrimSpace(newData.Email)

	fmt.Print("Masukkan password: ")
	newData.Password, _ = reader.ReadString('\n')
	newData.Password = strings.TrimSpace(newData.Password)

	// Panggil fungsi AddPegawai dari PegawaiModel untuk menyimpan data pegawai baru
	pegawai, err := pc.model.AddPegawai(newData)
	if err != nil {
		return fmt.Errorf("gagal menambahkan pegawai: %v", err)
	}

	fmt.Printf("Pegawai dengan ID %d berhasil ditambahkan.\n", pegawai.ID)
	return nil
}
