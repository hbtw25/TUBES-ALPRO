Berikut adalah `README.md` yang disesuaikan dengan struktur program Go yang telah kita kerjakan dan informasi pembagian tugas serta jadwal yang Anda berikan:

```markdown
# Sistem Pendaftaran Lomba Esport Antar Fakultas (TUBES-ALPRO)

Program ini adalah aplikasi berbasis terminal (CLI) yang dikembangkan menggunakan bahasa pemrograman Go. Aplikasi ini bertujuan untuk mengelola pendaftaran tim dalam sebuah turnamen esport antar fakultas, memungkinkan pengguna untuk menambah, melihat, mencari, memperbarui, menghapus, mengurutkan data tim, serta melihat statistik turnamen.

## Anggota Tim & Pembagian Tugas

### Harsya Brahmantyo Wibowo (Tyo)
1.  **Pengembangan Inti Program**:
    * Implementasi struktur data utama (`TimEsport`).
    * Fungsi-fungsi manajemen data inti (tambah, update, hapus tim secara fisik).
    * Implementasi algoritma pencarian (Sequential & Binary Search) dan pengurutan (Selection & Insertion Sort).
    * Integrasi seluruh komponen program.
2.  **Pengujian dan Debugging**:
    * Memastikan semua fitur berfungsi sesuai harapan.
    * Mengidentifikasi dan memperbaiki bug.
    * Membantu anggota tim lain dalam aspek teknis.
3.  **Kepemimpinan Tim**:
    * Mengoordinasikan alur pengembangan.
    * Membantu pemahaman kode program bagi anggota tim.

### Raditya Vihandika Bari Jabran (Raditya)
1.  **Pengembangan Fitur Statistik**:
    * Implementasi fungsi untuk menghitung statistik (rata-rata poin, poin tertinggi/terendah, jumlah tim).
    * Merancang dan mengimplementasikan tampilan untuk fitur statistik.
2.  **Pengembangan Antarmuka Pengguna (UI CLI)**:
    * Desain dan implementasi fungsi menu utama dan navigasi program.
    * Format tampilan data tim agar mudah dibaca.
    * Memastikan antarmuka pengguna intuitif dan ramah pengguna.
3.  **Dokumentasi Kode**:
    * Menambahkan komentar pada kode untuk menjelaskan fungsionalitas (IS: FS:).
    * Menyusun dan melengkapi file `README.md` ini.

### Muhamad Pradipa Dwi Pahlevi (Pradipa)
1.  **Dokumentasi dan Penyusunan Laporan Proyek**:
    * Menyusun deskripsi masalah yang diselesaikan oleh program.
    * Menjelaskan fitur-fitur program secara detail.
    * Menyiapkan tangkapan layar (screenshot) output program.
    * Mendokumentasikan pembagian tugas anggota tim.
2.  **Pengujian Fungsional**:
    * Melakukan pengujian program dari perspektif pengguna akhir.
    * Mencatat dan melaporkan bug atau masalah yang ditemukan.
    * Membantu dalam penyediaan data dummy untuk keperluan pengujian.
3.  **Pengembangan Komponen Pendukung**:
    * Implementasi fungsi-fungsi pendukung atau yang lebih sederhana seperti `tampilkanMenu()` dan `tampilkanSemuaTim()`.
    * Memberikan bantuan pada bagian kode lain yang lebih mudah dipahami.

---

## Jadwal dan Koordinasi Proyek

1.  **Minggu 1**:
    * **Tyo**: Mendesain struktur dasar program, implementasi struct `TimEsport`, dan fungsi-fungsi CRUD dasar.
    * **Raditya**: Mempelajari struktur program, mulai merancang UI CLI dan fungsi menu.
    * **Pradipa**: Mulai menyusun kerangka laporan, mempelajari alur program, dan mengidentifikasi kebutuhan data dummy.

2.  **Minggu 2**:
    * **Tyo**: Menyelesaikan implementasi algoritma pengurutan (Selection & Insertion Sort) dan pencarian (Sequential & Binary Search).
    * **Raditya**: Mengembangkan fungsi statistik dan menyelesaikan implementasi UI CLI.
    * **Pradipa**: Mulai membuat dokumentasi kode (komentar awal), melakukan pengujian fungsional awal pada fitur yang sudah ada.

3.  **Minggu 3-4**:
    * **Tyo**: Integrasi seluruh komponen, melakukan debugging menyeluruh, dan finalisasi core logic.
    * **Raditya**: Finalisasi dokumentasi kode (termasuk IS:FS: comments) dan penyusunan `README.md`.
    * **Pradipa**: Menyelesaikan laporan akhir proyek, melakukan pengujian akhir secara komprehensif, dan menyiapkan screenshot output.

---

## Deskripsi Program

Aplikasi ini memungkinkan panitia lomba untuk mengelola data tim esport yang mendaftar. Fitur-fitur yang disediakan mencakup CRUD (Create, Read, Update, Delete) data tim, perhitungan statistik performa tim, serta kemampuan untuk mencari dan mengurutkan data tim berdasarkan berbagai kriteria.

---

## Fitur Utama

1.  **Manajemen Data Tim**
    * **Tambah Tim**: Mendaftarkan tim baru dengan detail ID, Nama Tim, Fakultas, dan Poin.
    * **Lihat Data Tim**: Menampilkan seluruh tim yang terdaftar dalam format tabel.
    * **Cari Tim**:
        * Berdasarkan ID Tim (menggunakan Sequential Search atau Binary Search).
        * Berdasarkan Nama Tim (menggunakan Sequential Search atau Binary Search).
    * **Update Data Tim**: Mengubah informasi Nama Tim, Fakultas, atau Poin untuk tim yang sudah terdaftar.
    * **Hapus Tim**: Menghapus data tim dari sistem (penghapusan fisik).
2.  **Statistik Tim**
    * Menampilkan jumlah total tim yang terdaftar.
    * Menghitung dan menampilkan rata-rata poin semua tim.
    * Menampilkan tim dengan poin tertinggi (beserta nama tim).
    * Menampilkan tim dengan poin terendah (beserta nama tim).
3.  **Pengurutan Data Tim**
    * Pilihan algoritma pengurutan: **Selection Sort** atau **Insertion Sort**.
    * Kriteria pengurutan: ID Tim, Nama Tim, Fakultas, atau Poin.
    * Jenis pengurutan: Ascending (A-Z, 0-9) atau Descending (Z-A, 9-0).
    * Hasil pengurutan ditampilkan tanpa mengubah urutan data asli di sistem.

---

## Struktur Program

### Tipe Data Bentukan (Struct)

Data setiap tim disimpan menggunakan struct `TimEsport`:

```go
type TimEsport struct {
	ID       string
	Nama     string
	Fakultas string
	Poin     int
}
```

### Variabel Global

Program menggunakan variabel global untuk menyimpan data tim dan jumlah tim saat ini:

```go
const NMAX = 5 // Kapasitas maksimum tim (dapat disesuaikan dalam kode sumber)

var timData [NMAX]TimEsport // Array statis untuk menyimpan data tim.
var jumlahTim int = 0       // Jumlah tim yang saat ini tersimpan.
```

### Fungsi-fungsi Utama (Contoh berdasarkan kode yang ada)

* **Manajemen Data**: `tambahTim()`, `updateTim()`, `hapusTim()`, `tampilkanSemuaTim()`
* **Pencarian**:
    * `sequentialSearchByIDGlobal(id string) int`
    * `sequentialSearchByNamaGlobal(nama string) int`
    * `binarySearchByID(arr *[NMAX]TimEsport, n int, idCari string) int` (digunakan pada salinan data yang terurut)
    * `binarySearchByNama(arr *[NMAX]TimEsport, n int, namaCari string) int` (digunakan pada salinan data yang terurut)
* **Statistik**: `tampilkanStatistik()`
* **Pengurutan**:
    * `selectionSortByID(arr *[NMAX]TimEsport, n int, isAscending bool)`
    * `selectionSortByNama(arr *[NMAX]TimEsport, n int, isAscending bool)`
    * `selectionSortByFakultas(arr *[NMAX]TimEsport, n int, isAscending bool)`
    * `selectionSortByPoin(arr *[NMAX]TimEsport, n int, isAscending bool)`
    * `insertionSortByID(arr *[NMAX]TimEsport, n int, isAscending bool)`
    * `insertionSortByNama(arr *[NMAX]TimEsport, n int, isAscending bool)`
    * `insertionSortByFakultas(arr *[NMAX]TimEsport, n int, isAscending bool)`
    * `insertionSortByPoin(arr *[NMAX]TimEsport, n int, isAscending bool)`
* **Menu & Alur Program**: `main()`, `tampilkanMenu()`, `cariTim()`, `urutkanTim()`

---

## Algoritma yang Digunakan

### Pencarian Sekuensial (Sequential Search)
Digunakan untuk mencari tim berdasarkan ID atau Nama secara langsung pada data global (`timData`). Algoritma ini memeriksa setiap elemen dari awal hingga akhir sampai data ditemukan atau semua elemen telah diperiksa.

```go
// Contoh dari kode program:
func sequentialSearchByIDGlobal(id string) int {
	// IS: id = ID yang dicari. timData, jumlahTim adalah state global.
	// FS: Index i jika timData[i].ID == id, atau -1 jika tidak ditemukan.
	var i int = 0
	var found bool = false
	for i < jumlahTim && !found {
		if timData[i].ID == id {
			found = true
		} else {
			i++
		}
	}
	if found {
		return i
	}
	return -1
}
```

### Pencarian Biner (Binary Search)
Digunakan sebagai opsi pencarian tim berdasarkan ID atau Nama. **Penting**: Algoritma ini memerlukan data yang sudah terurut. Oleh karena itu, program membuat salinan data tim (`searchableTeamsArray`), mengurutkannya terlebih dahulu (menggunakan salah satu algoritma sorting yang ada), baru kemudian melakukan Binary Search pada salinan terurut tersebut.

```go
// Contoh dari kode program (pencarian berdasarkan ID pada array terurut):
func binarySearchByID(arr *[NMAX]TimEsport, n int, idCari string) int {
	// Precondition: arr[0...n-1] terurut berdasarkan ID.
	// IS: arr menunjuk ke array, n jumlah elemen, idCari adalah target. kr=0, kn=n-1, found=false.
	// FS: Index med jika arr[med].ID == idCari, atau -1 jika tidak.
	var kr, kn, med int
	kr = 0
	kn = n - 1
	var found bool = false
	for kr <= kn && !found {
		med = (kr + kn) / 2
		if arr[med].ID == idCari {
			found = true
		} else if arr[med].ID < idCari {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}
	// ... (return logic)
}
```

### Pengurutan Seleksi (Selection Sort)
Salah satu algoritma pengurutan yang disediakan. Algoritma ini secara berulang mencari elemen minimum (atau maksimum) dari bagian data yang belum terurut dan menukarnya dengan elemen pertama dari bagian yang belum terurut tersebut. Operasi pengurutan dilakukan pada salinan data (`timToSortArray`).

```go
// Contoh dari kode program (pengurutan berdasarkan Poin):
func selectionSortByPoin(arr *[NMAX]TimEsport, n int, isAscending bool) {
	// IS: arr menunjuk ke array, n jumlah elemen valid. isAscending menentukan urutan.
	// FS: n elemen pertama arr terurut berdasarkan Poin sesuai isAscending.
	var i, j, idx_min_max int
	var t TimEsport
	i = 1
	for i <= n-1 {
		idx_min_max = i - 1
		j = i
		for j < n {
			if isAscending {
				if arr[j].Poin < arr[idx_min_max].Poin { idx_min_max = j }
			} else {
				if arr[j].Poin > arr[idx_min_max].Poin { idx_min_max = j }
			}
			j = j + 1
		}
		t = arr[idx_min_max]
		arr[idx_min_max] = arr[i-1]
		arr[i-1] = t
		i = i + 1
	}
}
```

### Pengurutan Sisip (Insertion Sort)
Algoritma pengurutan alternatif yang juga disediakan. Algoritma ini membangun array akhir yang terurut satu elemen pada satu waktu. Setiap elemen baru disisipkan ke posisi yang tepat dalam sub-array yang sudah terurut. Operasi pengurutan dilakukan pada salinan data (`timToSortArray`).

```go
// Contoh dari kode program (pengurutan berdasarkan Poin):
func insertionSortByPoin(arr *[NMAX]TimEsport, n int, isAscending bool) {
	// IS: arr menunjuk ke array, n jumlah elemen valid. isAscending menentukan urutan.
	// FS: n elemen pertama arr terurut berdasarkan Poin sesuai isAscending.
	var i, j int
	var temp TimEsport
	for i = 1; i < n; i++ {
		temp = arr[i]
		j = i - 1
		if isAscending {
			for j >= 0 && arr[j].Poin > temp.Poin {
				arr[j+1] = arr[j]; j = j - 1
			}
		} else {
			for j >= 0 && arr[j].Poin < temp.Poin {
				arr[j+1] = arr[j]; j = j - 1
			}
		}
		arr[j+1] = temp
	}
}
```

---

## Cara Menggunakan Program

1.  **Kompilasi Program**:
    Buka terminal atau command prompt, navigasi ke direktori tempat file `main.go` disimpan, lalu jalankan perintah:
    ```bash
    go build main.go
    ```
    Ini akan menghasilkan sebuah file executable bernama `main` (atau `main.exe` di Windows).

2.  **Jalankan Program**:
    Eksekusi file yang telah dikompilasi:
    ```bash
    ./main
    ```
    (atau `.\main.exe` di Windows)

3.  **Interaksi dengan Menu**:
    Program akan menampilkan menu utama. Masukkan nomor pilihan yang sesuai untuk menggunakan fitur yang diinginkan.

---

## Batasan Program

* Program menggunakan array statis dengan kapasitas maksimum yang ditentukan oleh konstanta `NMAX` (saat ini `5`, namun dapat diubah dalam kode sumber untuk kapasitas yang berbeda).
* Program melakukan penghapusan data secara fisik (elemen digeser), bukan *soft delete*.
* Input pengguna diasumsikan sesuai dengan tipe data yang diminta (misalnya, angka untuk poin). Error handling untuk input yang tidak valid bersifat dasar.
* Tidak ada penggunaan `break` atau `continue` di dalam loop (sesuai batasan umum tugas Algoritma Pemrograman).

---
```
