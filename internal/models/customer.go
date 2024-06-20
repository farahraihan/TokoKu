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

func (cm *CustomerModel) UpdateCustomer(id uint, updatedData Customer) (Customer, error) {
    var existingData Customer
    // Cari data pelanggan berdasarkan ID
    err := cm.db.First(&existingData, id).Error
    if err != nil {
        return Customer{}, err
    }

    // Update data yang ditemukan
    existingData.Username = updatedData.Username
    existingData.Nama = updatedData.Nama
    existingData.Email = updatedData.Email
    existingData.NoTelp = updatedData.NoTelp

    // Simpan data yang telah diperbarui kembali ke database
    err = cm.db.Save(&existingData).Error
    if err != nil {
        return Customer{}, err
    }

    return existingData, nil
}
func (cm *CustomerModel) DeleteCustomer(id uint) error {
	var customer Customer
	err := cm.db.First(&customer, id).Error
	if err != nil {
		return err
	}
	err = cm.db.Delete(&customer).Error
	if err != nil {
		return err
	}
	return nil
}