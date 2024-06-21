# TokoKu
Tokoku adalah sebuah aplikasi yang digunakan untuk mencatat kegiatan transaksi di sebuah toko secara online.

## Project Root
- Project Root
  - configs /
  - internal /
    - controllers /
    - models /
  - readme.md
  - .gitignore
  - go.mod
  - go.sum
  - main.go


### `configs`

Memuat code yang bertujuan untuk konfigurasi ke database

### `internal`

Berisi rangkaian utama code yang digunakan dalam project. 

### `controllers`

Berisi code yang mencakup fungsi fitur yang diakomodasi oleh program. Dalam folder ini, setiap fitur terbagi dalam beberapa file yaitu : pegawai.go, barang.go, customer.go, transaksi.go, detail_transaksi.go.

### `models`

Berisi code yang bertugas untuk berhubungan dengan database. Code yang diletakkan dalam folder ini sangat spesifik digunakan untuk bersinggungan dengan database, segala sesuatau yang berhubungan dengan ORM dan database dilakukan pada folder ini. Folder ini terbagi lagi ke dalam beberapa file yaitu : pegawai.go, barang.go, customer.go, transaksi.go, detail_transaksi.go.

## Fitur
Berikut fitur prioritas untuk memenuhi kriteria MVP (Minimum Viable Product):
- Terdapat sebuah akun admin (username : admin, password : admin)
- Admin bertugas untuk mendaftarkan pegawai lain agar dapat login ke dalam sistem
- Admin DAPAT MENGHAPUS DATA APAPUN PADA SISTEM
- Admin dapat melakukan hal apapun yang bisa dilakukan oleh pegawai biasa.
- Pegawai dapat melakukan login kedalam sistem.
- Pegawai dapat menambahkan barang baru ke dalam sistem
- Pegawai dapat melakukan edit informasi barang
- Pegawai dapat melakukan update stok barang
- Pegawai dapat menambahkan daftar customer
- Pegawai dapat membuat nota transaksi untuk customer

- Catatan : Untuk setiap barang yang dibuat, wajib disertai dengan siapa yang menambahkan barang tersebut. Barang yang sudah ditambahkan dapat diakses oleh semua pegawai.

Transaksi memiliki kriteria sebagai berikut:

- Wajib mengandung tanggal pembuatan transaksi (biasanya tanggal hari ini)
- Wajib terdapat identitas pegawai yang membuat transaksi
- Dalam 1 transaksi bisa memiliki banyak barang
- Stok barang yang dituliskan dalam transaksi akan mengurangi stok barang yang tersedia
- PEGAWAI TIDAK BISA MENGUBAH TRANSAKSI YANG TELAH DIBUAT
- Terdapat fitur untuk melihat daftar barang yang ada dalam sistem
