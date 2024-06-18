package models

type Barang struct {
	Nama  string `gorm:"primarykey"`
	Harga int
	Stok  int
}
