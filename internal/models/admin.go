package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string
	Nama     string
	Gender   Gender
	NoTelp   string
	Password string
	Email    string
	Pegawais []Pegawai `gorm:"foreignKey:AdminID"`
}

type AdminModel struct {
	db *gorm.DB
}

func NewAdminModel(connection *gorm.DB) *AdminModel {
	return &AdminModel{
		db: connection,
	}
}

// login admin
