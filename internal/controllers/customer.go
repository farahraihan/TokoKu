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

func (cc *CustomerController) GetCustomer() {
	customers, err := cc.model.GetCustomer()
	if err != nil {
		fmt.Println("Error:", err)
		return
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
}
