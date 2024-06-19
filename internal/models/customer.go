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
func (cm *CustomerModel) AddCustomer(newData Customer) (Customer, error) {
	err := cm.db.Create(&newData).Error
	if err != nil {
		return Customer{}, err
	}
	return newData, nil
}