package main

import (
	"fmt"
	"tokoku/configs"
	"tokoku/internal/controllers"
	"tokoku/internal/models"
)

func main() {

	// Koneksi ke database
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database:", err.Error())
		return
	}

	connection.AutoMigrate(&models.Pegawai{}, &models.Customer{}, &models.Transaksi{}, &models.Barang{}, &models.DetailTransaksi{})

	pm := models.NewPegawaiModel(connection)
	pc := controllers.NewPegawaiController(pm)

	cm := models.NewCustomerModel(connection)
	cc := controllers.NewCustomerController(cm)

	bm := models.NewBarangModel(connection)
	bc := controllers.NewBarangController(bm)

	dm := models.NewDetailTransaksiModel(connection, bm)
	dc := controllers.NewDetailTransaksiController(dm, bc)

	tm := models.NewTransaksiModel(connection)
	tc := controllers.NewTransaksiController(tm, bm, dc)

	fmt.Print("\nSELAMAT DATANG DI LAMAN TOKOKU ! ^_^\n")
	fmt.Println("____________________________________")

	var inputMenu int

	for inputMenu != 1000 {
		fmt.Println("Silahkan pilih menu :")
		fmt.Println("1. Login ")
		fmt.Println("2. Keluar ")
		fmt.Print("Masukkan input: ")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:

			pegawai, err := pc.Login()
			if err != nil {
				fmt.Println("Terjadi error pada saat login, error: ", err.Error())
				return
			}
			username := pegawai.Username

			if username == "admin" {
				fmt.Print("\nTerdeteksi sebagai admin\n")

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
								fmt.Println("____________________________________")
								err = pc.AddPegawai()
								if err != nil {
									fmt.Printf("Error: %v\n", err)
								} else {
									fmt.Println("Pegawai berhasil ditambahkan.")
								}

							case 2:
								fmt.Println("Edit data pegawai")
								fmt.Println("____________________________________")
								err = pc.UpdatePegawai()
								if err != nil {
									fmt.Printf("Error: %v\n", err)
									return
								} else {
									fmt.Println("Berhasil edit data pegawai")
								}

							case 3:
								fmt.Println("Hapus data pegawai")
								fmt.Println("____________________________________")
								_, err := pc.DeletePegawai()
								if err != nil {
									fmt.Printf("Error: %v\n", err)
									return
								} else {
									fmt.Println("Pegawai berhasil dihapus.")
								}

							case 4:
								fmt.Println("Lihat data pegawai")
								fmt.Println("____________________________________")
								err = pc.GetPegawai()
								if err != nil {
									fmt.Printf("Error: %v\n", err)
									return
								} else {
									fmt.Println("Berhasil menampilkan data pegawai")
								}
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
								fmt.Println("____________________________________")
								_, err := cc.AddCustomer(pegawai.ID)
								if err != nil {
									fmt.Printf("Error: %v\n", err)
									return
								} else {
									fmt.Println("Customer berhasil ditambahkan.")
								}

							case 2:
								fmt.Println("Edit data customer")
								fmt.Println("____________________________________")
								_, err := cc.UpdateCustomer(pegawai.ID)
								if err != nil {
									fmt.Printf("Error: %v\n", err)
									return
								} else {
									fmt.Println("Customer berhasil diedit")
								}

							case 3:
								fmt.Println("Hapus data customer")
								fmt.Println("____________________________________")
								_, err := cc.DeleteCustomer()
								if err != nil {
									fmt.Printf("Error: %v\n", err)
									return
								} else {
									fmt.Println("Customer berhasil dihapus")
								}

							case 4:
								fmt.Println("Lihat data customer")
								fmt.Println("____________________________________")
								err = cc.GetCustomer()
								if err != nil {
									fmt.Printf("Error: %v\n", err)
									return
								} else {
									fmt.Println("Berhasil menampilkan data customer")
								}
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
							fmt.Println("2. Hapus data transaksi")
							fmt.Println("3. Cetak transaksi")
							fmt.Println("4. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								fmt.Println("Tambah data transaksi")
								fmt.Println("____________________________________")
								// transaksiID, err = tc.AddTransaksi(pegawai.ID)
								// if err != nil {
								// 	fmt.Printf("Error: %v\n", err)
								// 	return
								// } else {
								// 	fmt.Println("Transaksi berhasil ditambahkan")
								// 	err = dc.AddDetailTransaksi(transaksiID)
								// 	if err != nil {
								// 		fmt.Printf("Error: %v\n", err)
								// 		return
								// 	}
								// 	fmt.Println("Detail transaksi berhasil ditambahkan.")
								// }

							case 2:
								fmt.Println("Hapus data transaksi")
								fmt.Println("____________________________________")
								_, err := tc.DeleteTransaksi()
								if err != nil {
									fmt.Printf("Error: %v\n", err)
									return
								} else {
									fmt.Println("Transaksi berhasil dihapus")
								}

							case 3:
								fmt.Println("Cetak nota transaksi")
								fmt.Println("____________________________________")
								err := tc.PrintNotaTransaksi()
								if err != nil {
									fmt.Printf("Error: %v\n", err)
									return
								} else {
									fmt.Println("Berhasil menampilkan nota transaksi")
								}
							case 4:
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
							fmt.Println("3. Tambah stock barang")
							fmt.Println("4. Hapus data barang")
							fmt.Println("5. Lihat data barang")
							fmt.Println("6. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								fmt.Println("Tambah data barang")
								_, err := bc.AddBarang(1)
								if err != nil {
									fmt.Printf("Error adding barang: %v\n", err)
									return
								}
								fmt.Println("Barang berhasil ditambahkan.")
							case 2:
								fmt.Println("Edit data barang")
								_, err := bc.UpdateBarang()
								if err != nil {
									fmt.Printf("Error Updating Barang: %v\n", err)
									return
								}
								fmt.Println("Barang berhasil diupdate.")
							case 3:
								fmt.Println("Tambah stock barang")
								// Contoh penggunaan UpdateStock
								err = bc.UpdateStock()
								if err != nil {
									fmt.Println("Error:", err)
								}
							case 4:
								fmt.Println("Hapus data barang")
								_, err := bc.DeleteBarang()
								if err != nil {
									fmt.Printf("Error Delete Barang: %v\n", err)
									return
								}
								fmt.Println("Barang berhasil dihapus.")
							case 5:
								fmt.Println("Lihat data barang")
								// bc.GetBarang()
							case 6:
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

			} else {
				fmt.Print("\nTerdeteksi sebagai pegawai\n")

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
							fmt.Println("2. Cetak transaksi")
							fmt.Println("3. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								fmt.Println("Tambah data transaksi")
								// Add Transaksi dengan contoh id pegawai/admin
								transaksiID, err := tc.AddTransaksi(1)
								if err != nil {
									fmt.Printf("Error adding transaksi: %v\n", err)
									return
								}

								// Memanggil fungsi AddDetailTransaksi dari controller
								err = dc.AddDetailTransaksi(transaksiID)
								if err != nil {
									fmt.Printf("Error adding detail transaksi: %v\n", err)
									return
								}
								fmt.Println("Detail transaksi berhasil ditambahkan.")
							case 2:
								err := tc.PrintNotaTransaksi()
								if err != nil {
									fmt.Printf("Error: %v\n", err)
								}
							case 3:
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
							fmt.Println("3. Tambah stock barang")
							fmt.Println("4. Hapus data barang")
							fmt.Println("5. Lihat data barang")
							fmt.Println("6. Keluar")
							fmt.Print("Masukkan input: ")
							fmt.Scanln(&inputMenu)
							switch inputMenu {
							case 1:
								fmt.Println("Tambah data barang")
								_, err := bc.AddBarang(1)
								if err != nil {
									fmt.Printf("Error adding barang: %v\n", err)
									return
								}
								fmt.Println("Barang berhasil ditambahkan.")
							case 2:
								fmt.Println("Edit data barang")
								_, err := bc.UpdateBarang()
								if err != nil {
									fmt.Printf("Error Updating Barang: %v\n", err)
									return
								}
								fmt.Println("Barang berhasil diupdate.")

							case 3:
								fmt.Println("Tambah stock barang")
								// Contoh penggunaan UpdateStock
								err = bc.UpdateStock()
								if err != nil {
									fmt.Println("Error:", err)
								}

							case 4:
								fmt.Println("Hapus data barang")
								_, err := bc.DeleteBarang()
								if err != nil {
									fmt.Printf("Error Delete Barang: %v\n", err)
									return
								}
								fmt.Println("Barang berhasil dihapus.")
							case 5:
								fmt.Println("Lihat data barang")
								// bc.GetBarang()
							case 6:
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

		case 2:
			fmt.Println("TERIMAKASIH ^_^ !")
			inputMenu = 1000

		default:
			fmt.Print("\nInput tidak valid ! pastikan hanya menulis angka\n")

		}
	}
}
