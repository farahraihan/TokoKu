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

func (cc *CustomerController) GetCustomer() error {
	customers, err := cc.model.GetCustomer()
	if err != nil {
		return fmt.Errorf("failed to get customers: %w", err)
	}

	for _, customer := range customers {
		fmt.Println("----------------------------------")
		fmt.Printf("ID: %d\n", customer.ID)
		fmt.Printf("PegawaiID: %d\n", customer.PegawaiID)
		fmt.Printf("Username: %s\n", customer.Username)
		fmt.Printf("Nama: %s\n", customer.Nama)
		fmt.Printf("Email: %s\n", customer.Email)
		fmt.Printf("NoTelp: %s\n", customer.NoTelp)
		fmt.Println()
	}

	return nil
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

func (cc *CustomerController) UpdateCustomer() (bool, error) {
	var id uint
	var updatedData models.Customer
	fmt.Print("Masukkan ID customer yang ingin diperbarui: ")
	fmt.Scanln(&id)
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&updatedData.Username)
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&updatedData.Nama)
	fmt.Print("Masukkan Email: ")
	// Membersihkan newline dari buffer
	fmt.Scanln()

	fmt.Scanln(&updatedData.Email)
	fmt.Print("Masukkan nomor telepon: ")
	fmt.Scanln(&updatedData.NoTelp)

	_, err := cc.model.UpdateCustomer(id, updatedData)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (cc *CustomerController) DeleteCustomer() (bool, error) {
	var deleteData models.Customer
	fmt.Print("Masukkan ID customer yang ingin dihapus: ")
	fmt.Scanln(&deleteData.ID)
	err := cc.model.DeleteCustomer(deleteData.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
