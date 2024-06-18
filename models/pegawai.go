// models/pegawai_model.go

package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

// Struktur model untuk pegawai
type Pegawai struct {
	Username string
	Nama     string
	Kelamin  string
	NoTelp   int
	Email    string
	Password string
}

// Fungsi untuk melakukan login pegawai
func LoginPegawai(username, password string) bool {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return false
	}
	defer db.Close()

	// Query untuk memeriksa keberadaan pegawai dengan username dan password tertentu
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM pegawai WHERE username = ? AND password = ?", username, password).Scan(&count)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return false
	}

	return count == 1
}

// Fungsi lain untuk mengelola data pegawai, seperti menambah, mengubah, menghapus, atau mengambil data
// Implementasikan sesuai kebutuhan aplikasi Anda
