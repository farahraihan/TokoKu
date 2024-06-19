package controllers

import (
	"fmt"
	"tokoku/internal/models"
)

type CustomerController struct {
	model *models.CustomerModel
}

func NewCustomerController(m *models.CustomerModel) *CustomerController {
	return &CustomerController{
		model: m,
	}
}

func (cc *CustomerController) AddCustomer(id uint) (bool, error) {
	var newData models.Customer
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&newData.Username)
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&newData.Nama)
	fmt.Print("Masukkan Email: ")
	// Membersihkan newline dari buffer
	fmt.Scanln()
	
	fmt.Scanln(&newData.Email)
	fmt.Print("Masukkan nomor telepon: ")
	fmt.Scanln(&newData.NoTelp)
	newData.PegawaiID = id
	_, err := cc.model.AddCustomer(newData)
	if err != nil {
		return false, err
	}

	return true, nil
}