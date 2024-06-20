package main

import (
	"fmt"
	"tokoku/configs"
	"tokoku/internal/models"
	// "tokoku/internal/controllers"
)

func main() {

	// Koneksi ke database
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database:", err.Error())
		return
	}

	connection.AutoMigrate(&models.Admin{}, &models.Transaksi{}, &models.Barang{}, &models.DetailTransaksi{})

	// am := models.NewAdminModel(connection)
	// ac := controllers.NewAdminController(am)

	// pm := models.NewPegawaiModel(connection)
	// pc := controllers.NewPegawaiController(pm)

	// cm := models.NewCustomerModel(connection)
	// cc := controllers.NewCustomerController(cm)

	

	// bm := models.NewBarangModel(connection)
	// bc := controllers.NewBarangController(bm)

	// dm := models.NewDetailTransaksiModel(connection,bm)
	// dc := controllers.NewDetailTransaksiController(dm,bc)

	// tm := models.NewTransaksiModel(connection)
	// tc := controllers.NewTransaksiController(tm,dc)

	fmt.Print("\nSELAMAT DATANG DI LAMAN TOKOKU ! ^_^\n")
	fmt.Println("____________________________________")

	var inputMenu int

	for inputMenu != 1000 {
		fmt.Println("Silahkan pilih menu :")
		fmt.Println("1. Login sebagai admin")
		fmt.Println("2. Login sebagai pegawai")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan input: ")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:
			// manggil function ac.login
			var isLogin = true

			if isLogin {
				for inputMenu != 1000 {
					fmt.Println("____________________________________")
					fmt.Println("Silahkan pilih menu :")
					fmt.Println("1. DATA PEGAWAI")
					fmt.Println("2. DATA CUSTOMER")
					fmt.Println("3. DATA TRANSAKSI")
					fmt.Println("4. DATA BARANG")
					fmt.Println("5. Keluar")
					fmt.Print("Masukkan input: ")
					fmt.Scanln(&inputMenu)

					switch inputMenu {
					case 1:
						for inputMenu != 1000 {
							fmt.Println("____________________________________")
							fmt.Println("Pilih aksi : ")
							fmt.Println("1. Tambah data pegawai")
							fmt.Println("2. Edit data pegawai")
							fmt.Println("3. Hapus data pegawai")
							fmt.Println("4. Lihat data pegawai")
							fmt.Println("5. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)

							switch inputMenu {
							case 1:
								fmt.Println("Tambah data pegawai")
								// pc.AddPegawai()
							case 2:
								fmt.Println("Edit data pegawai")
								// pc.UpdatePegawai()
							case 3:
								fmt.Println("Hapus data pegawai")
								// pc.DeletePegawai()
							case 4:
								fmt.Println("Lihat data pegawai")
								// pc.GetPegawai()
							case 5:
								fmt.Println("TERIMAKASIH ^_^ !")
								inputMenu = 1000
							default:
								fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")
							}
						}

					case 2:
						for inputMenu != 1000 {
							fmt.Println("____________________________________")
							fmt.Println("Pilih aksi : ")
							fmt.Println("1. Tambah data customer")
							fmt.Println("2. Edit data customer")
							fmt.Println("3. Hapus data customer")
							fmt.Println("4. Lihat data customer")
							fmt.Println("5. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								fmt.Println("Tambah data customer")
								// cc.AddCustomer()
							case 2:
								fmt.Println("Edit data customer")
								// cc.UpdateCustomer()
							case 3:
								fmt.Println("Hapus data customer")
								// cc.DeleteCustomer()
							case 4:
								fmt.Println("Lihat data customer")
								// cc.GetCustomer()
							case 5:
								fmt.Println("TERIMAKASIH ^_^ !")
								inputMenu = 1000
							default:
								fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")
							}
						}

					case 3:
						for inputMenu != 1000 {
							fmt.Println("____________________________________")
							fmt.Println("Pilih aksi : ")
							fmt.Println("1. Tambah data transaksi")
							fmt.Println("2. Edit data transaksi")
							fmt.Println("3. Hapus data transaksi")
							fmt.Println("4. Lihat data transaksi")
							fmt.Println("5. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								fmt.Println("Tambah data transaksi")
								// tc.AddTransaksi()
							case 2:
								fmt.Println("Edit data transaksi")
								// tc.UpdateTransaksi()
							case 3:
								fmt.Println("Hapus data transaksi")
								// tc.DeleteTransaksi()
							case 4:
								fmt.Println("Lihat data transaksi")
								// tc.GetTransaksi()
							case 5:
								fmt.Println("TERIMAKASIH ^_^ !")
								inputMenu = 1000
							default:
								fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")
							}
						}

					case 4:
						for inputMenu != 1000 {
							fmt.Println("____________________________________")
							fmt.Println("Pilih aksi : ")
							fmt.Println("1. Tambah data barang")
							fmt.Println("2. Edit data barang")
							fmt.Println("3. Hapus data barang")
							fmt.Println("4. Lihat data barang")
							fmt.Println("5. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								fmt.Println("Tambah data barang")
								// bc.AddBarang()
							case 2:
								fmt.Println("Edit data barang")
								// bc.UpdateBarang()
							case 3:
								fmt.Println("Hapus data barang")
								// bc.DeleteBarang()
							case 4:
								fmt.Println("Lihat data barang")
								// bc.GetBarang()
							case 5:
								fmt.Println("TERIMAKASIH ^_^ !")
								inputMenu = 1000
							default:
								fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")
							}
						}

					case 5:
						fmt.Println("TERIMAKASIH ^_^ !")
						inputMenu = 1000

					default:
						fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")

					}
				}
			}

		case 2:
			// Manggil function pc.Login()

			var isLogin = true

			if isLogin {
				for inputMenu != 1000 {
					fmt.Println("____________________________________")
					fmt.Println("Silahkan pilih menu :")
					fmt.Println("1. DATA CUSTOMER")
					fmt.Println("2. DATA TRANSAKSI")
					fmt.Println("3. DATA BARANG")
					fmt.Println("4. Keluar")
					fmt.Print("Masukkan input: ")
					fmt.Scanln(&inputMenu)

					switch inputMenu {
					case 1:
						for inputMenu != 1000 {
							fmt.Println("____________________________________")
							fmt.Println("Pilih aksi : ")
							fmt.Println("1. Tambah data customer")
							fmt.Println("2. Edit data customer")
							fmt.Println("3. Hapus data customer")
							fmt.Println("4. Lihat data customer")
							fmt.Println("5. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								fmt.Println("Tambah data customer")
								// cc.AddCustomer()
							case 2:
								fmt.Println("Edit data customer")
								// cc.UpdateCustomer()
							case 3:
								fmt.Println("Hapus data customer")
								// cc.DeleteCustomer()
							case 4:
								fmt.Println("Lihat data customer")
								// cc.GetCustomer()
							case 5:
								fmt.Println("TERIMAKASIH ^_^ !")
								inputMenu = 1000
							default:
								fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")
							}
						}

					case 2:
						for inputMenu != 1000 {
							fmt.Println("____________________________________")
							fmt.Println("Pilih aksi : ")
							fmt.Println("1. Tambah data transaksi")
							fmt.Println("2. Edit data transaksi")
							fmt.Println("3. Hapus data transaksi")
							fmt.Println("4. Lihat data transaksi")
							fmt.Println("5. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								// tc.AddTransaksi(1)
							case 2:
								fmt.Println("Edit data transaksi")
								// tc.UpdateTransaksi()
							case 3:
								fmt.Println("Hapus data transaksi")
								// tc.DeleteTransaksi()
							case 4:
								fmt.Println("Lihat data transaksi")
								// tc.GetTransaksi()
							case 5:
								fmt.Println("TERIMAKASIH ^_^ !")
								inputMenu = 1000
							default:
								fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")
							}
						}

					case 3:
						for inputMenu != 1000 {
							fmt.Println("____________________________________")
							fmt.Println("Pilih aksi : ")
							fmt.Println("1. Tambah data barang")
							fmt.Println("2. Edit data barang")
							fmt.Println("3. Hapus data barang")
							fmt.Println("4. Lihat data barang")
							fmt.Println("5. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								fmt.Println("Tambah data barang")
								// bc.AddBarang()
							case 2:
								fmt.Println("Edit data barang")
								// bc.UpdateBarang()
							case 3:
								fmt.Println("Hapus data barang")
								// bc.DeleteBarang()
							case 4:
								fmt.Println("Lihat data barang")
								// bc.GetBarang()
							case 5:
								fmt.Println("TERIMAKASIH ^_^ !")
								inputMenu = 1000
							default:
								fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")
							}
						}

					case 4:
						fmt.Println("TERIMAKASIH ^_^ !")
						inputMenu = 1000

					default:
						fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")

					}

				}
			}

		case 3:
			fmt.Println("TERIMAKASIH ^_^ !")
			inputMenu = 1000

		default:
			fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")

		}
	}
}
