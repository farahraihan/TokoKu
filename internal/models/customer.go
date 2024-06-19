package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	PegawaiID  uint
	Username   string
	Nama       string
	Email      string
	NoTelp     string
	Transaksis []Transaksi `gorm:"foreignKey:CustomerID"`
}

type CustomerModel struct {
	db *gorm.DB
}

func NewCustomerModel(connection *gorm.DB) *CustomerModel {
	return &CustomerModel{
		db: connection,
	}
}

// GetCustomer
func (cm *CustomerModel) GetCustomer() ([]Customer, error) {
	var customers []Customer
	err := cm.db.Find(&customers).Error

	if err != nil {
		return nil, err
	}

	return customers, nil
}
