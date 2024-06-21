package controllers

import (
	"bufio"
	"errors"
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

func (pc *PegawaiController) GetPegawai() error {
	pegawais, err := pc.model.GetPegawai()
	if err != nil {
		return fmt.Errorf("failed to get pegawai: %w", err)
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

	return nil
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

func (pc *PegawaiController) UpdatePegawai() error {
	var updateData models.Pegawai
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pegawai yang ingin diupdate: ")
	var id uint
	fmt.Scanln(&id)

	fmt.Print("Masukkan username baru: ")
	updateData.Username, _ = reader.ReadString('\n')
	updateData.Username = strings.TrimSpace(updateData.Username)

	fmt.Print("Masukkan nama baru: ")
	updateData.Nama, _ = reader.ReadString('\n')
	updateData.Nama = strings.TrimSpace(updateData.Nama)

	fmt.Print("Masukkan gender baru (Pria/Wanita): ")
	gender, _ := reader.ReadString('\n')
	updateData.Gender = models.Gender(strings.TrimSpace(gender))

	if updateData.Gender != models.Pria && updateData.Gender != models.Wanita {
		return fmt.Errorf("gender tidak valid. Harus 'Pria' atau 'Wanita'")
	}

	fmt.Print("Masukkan nomor telepon baru: ")
	updateData.NoTelp, _ = reader.ReadString('\n')
	updateData.NoTelp = strings.TrimSpace(updateData.NoTelp)

	fmt.Print("Masukkan email baru: ")
	updateData.Email, _ = reader.ReadString('\n')
	updateData.Email = strings.TrimSpace(updateData.Email)

	fmt.Print("Masukkan password baru: ")
	updateData.Password, _ = reader.ReadString('\n')
	updateData.Password = strings.TrimSpace(updateData.Password)

	// Panggil fungsi UpdatePegawai dari PegawaiModel untuk mengupdate data pegawai
	_, err := pc.model.UpdatePegawai(id, updateData)
	if err != nil {
		return fmt.Errorf("gagal mengupdate pegawai: %v", err)
	}

	fmt.Printf("Pegawai dengan ID %d berhasil diupdate.\n", id)
	return nil
}
