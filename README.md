# TUBES-ALPRO


## Pembagian Tugas

### Tyo 
1. **Pengembangan Core Program**:
   - Implementasi struktur data utama (struct TimEsport)
   - Fungsi-fungsi utama manajemen data (tambah, edit, hapus tim)
   - Implementasi algoritma pencarian dan pengurutan
   - Integrasi seluruh komponen program

2. **Pengujian dan Debugging**:
   - Memastikan semua fitur berjalan dengan baik
   - Mengatasi bug yang muncul
   - Membantu anggota lain jika mengalami kesulitan teknis

3. **Mengarahkan tim**:
   - Mengoordinasi pengembangan program
   - Membantu anggota tim dalam memahami kode program

### Raditya 
1. **Pengembangan Fitur Statistik**:
   - Implementasi fungsi perhitungan statistik (rata-rata poin, poin tertinggi/terendah)
   - Implementasi tampilan statistik tim

2. **Pengembangan UI CLI Program**:
   - Fungsi menu dan navigasi program
   - Format tampilan data tim
   - Pembuatan antarmuka yang user-friendly

3. **Dokumentasi Kode**:
   - Menambahkan komentar pada kode untuk memperjelas fungsi
   - Menyiapkan file README

### Pradipa 
1. **Dokumentasi dan Laporan**:
   - Menyusun laporan deskripsi masalah
   - Menjelaskan fitur-fitur program
   - Menyiapkan screenshot output program
   - Menjelaskan pembagian tugas

2. **Testing**:
   - Melakukan pengujian program dari sisi pengguna
   - Mencatat bug dan masalah yang ditemukan
   - Membantu dalam pengumpulan data dummy untuk pengujian

3. **Pekerjaan yang Lebih Sederhana**:
   - Implementasi fungsi sederhana seperti tampilkanMenu() dan tampilkanSemuaTim()
   - Bantuan dalam bagian lain yang lebih mudah dipahami

## Jadwal dan Koordinasi

1. **Minggu 1**: 
   - Tyo: Membuat struktur dasar program dan implementasi fungsi-fungsi utama
   - Raditya: Mempelajari struktur dasar dan mulai bekerja pada UI
   - Pradipa: Mulai menyusun dokumen laporan dan mempelajari program

2. **Minggu 2**:
   - Tyo: Menyelesaikan implementasi algoritma pengurutan dan pencarian
   - Raditya: Mengembangkan fungsi statistik dan menyelesaikan UI
   - Pradipa: Mulai membuat dokumentasi kode dan melakukan pengujian awal

3. **Minggu 3-4**:
   - Tyo: Integrasi seluruh komponen dan debugging
   - Raditya: Penyelesaian dokumentasi kode dan README
   - Pradipa: Menyelesaikan laporan dan melakukan pengujian akhir


# Sistem Pendaftaran Lomba Esport Antar Fakultas

Program ini adalah aplikasi berbasis terminal untuk mengelola pendaftaran tim lomba esport antar fakultas, yang diimplementasikan menggunakan bahasa pemrograman Go.

## Deskripsi Program

Program ini dirancang untuk memudahkan pengelolaan data tim esport dalam sebuah lomba antar fakultas, termasuk pencatatan data tim, perhitungan statistik, pencarian data, dan pengurutan data tim berdasarkan berbagai kriteria.

## Fitur Utama

1. **Manajemen Data Tim**
   - Tambah tim baru dengan detail ID, nama, fakultas, dan poin
   - Tampilkan daftar semua tim yang terdaftar
   - Cari tim berdasarkan ID atau nama
   - Update data tim (nama, fakultas, poin)
   - Hapus tim dari sistem

2. **Statistik Tim**
   - Jumlah tim yang terdaftar
   - Rata-rata poin dari semua tim
   - Tim dengan poin tertinggi
   - Tim dengan poin terendah

3. **Pengurutan Data**
   - Pengurutan menggunakan algoritma Selection Sort
   - Pengurutan menggunakan algoritma Insertion Sort
   - Opsi pengurutan berdasarkan ID, nama, fakultas, atau poin
   - Dukungan pengurutan ascending (menaik) dan descending (menurun)

4. **Pencarian Data**
   - Pencarian tim berdasarkan ID menggunakan sequential search
   - Pencarian tim berdasarkan nama menggunakan sequential search

## Struktur Program

### Tipe Data

Program menggunakan tipe bentukan (struct) untuk menyimpan data tim:

```go
type TimEsport struct {
	ID        string
	Nama      string
	Fakultas  string
	Poin      int
	IsDeleted bool // Flag untuk penanda data yang sudah dihapus
}
```

### Variabel Global

```go
var timData [100]TimEsport // Array statis dengan kapasitas 100 tim
var jumlahTim int = 0      // Jumlah tim yang tersimpan
```

### Fungsi-fungsi Utama

1. **Fungsi Menu dan Navigasi**
   - `tampilkanMenu()` - Menampilkan menu utama program
   - `main()` - Fungsi utama yang mengatur alur program

2. **Fungsi Manajemen Data**
   - `tambahTim()` - Menambah data tim baru
   - `tampilkanSemuaTim()` - Menampilkan daftar semua tim
   - `tampilkanDetailTim()` - Menampilkan detail tim tertentu
   - `updateTim()` - Memperbarui data tim
   - `hapusTim()` - Menghapus data tim (menggunakan flag)

3. **Fungsi Pencarian**
   - `cariTim()` - Fungsi menu untuk pencarian tim
   - `cariTimByID()` - Mencari tim berdasarkan ID (sequential search)
   - `cariTimByNama()` - Mencari tim berdasarkan nama (sequential search)

4. **Fungsi Statistik**
   - `tampilkanStatistik()` - Menampilkan statistik tim (jumlah, rata-rata poin, poin tertinggi, poin terendah)

5. **Fungsi Pengurutan**
   - **Selection Sort**
     - `selectionSortByID()` - Pengurutan berdasarkan ID
     - `selectionSortByNama()` - Pengurutan berdasarkan nama
     - `selectionSortByFakultas()` - Pengurutan berdasarkan fakultas
     - `selectionSortByPoin()` - Pengurutan berdasarkan poin

   - **Insertion Sort**
     - `insertionSortByID()` - Pengurutan berdasarkan ID
     - `insertionSortByNama()` - Pengurutan berdasarkan nama
     - `insertionSortByFakultas()` - Pengurutan berdasarkan fakultas
     - `insertionSortByPoin()` - Pengurutan berdasarkan poin

## Algoritma Pencarian dan Pengurutan

### Sequential Search

Program menggunakan algoritma sequential search untuk mencari tim berdasarkan ID atau nama. Algoritma ini mencari secara berurutan dari awal hingga akhir array sampai menemukan data yang dicari.

```go
func cariTimByID(id string) int {
	for i := 0; i < jumlahTim; i++ {
		if timData[i].ID == id && !timData[i].IsDeleted {
			return i
		}
	}
	return -1
}
```

### Selection Sort

Algoritma selection sort diimplementasikan untuk mengurutkan data tim berdasarkan berbagai kriteria. Algoritma ini secara iteratif mencari elemen terkecil (atau terbesar) dari bagian array yang belum diurutkan dan memindahkannya ke posisi yang tepat.

```go
func selectionSortByPoin(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if isAscending {
				if arr[j].Poin < arr[minIdx].Poin {
					minIdx = j
				}
			} else {
				if arr[j].Poin > arr[minIdx].Poin {
					minIdx = j
				}
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}
```

### Insertion Sort

Algoritma insertion sort juga diimplementasikan untuk pengurutan data tim. Algoritma ini membangun array terurut satu elemen pada satu waktu dengan menempatkan setiap elemen baru pada posisi yang tepat dalam subarrai yang telah diurutkan.

```go
func insertionSortByPoin(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		
		if isAscending {
			for j >= 0 && arr[j].Poin > key.Poin {
				arr[j+1] = arr[j]
				j--
			}
		} else {
			for j >= 0 && arr[j].Poin < key.Poin {
				arr[j+1] = arr[j]
				j--
			}
		}
		
		arr[j+1] = key
	}
}
```

## Cara Menggunakan Program

1. Kompilasi program dengan perintah:
   ```
   go build main.go
   ```

2. Jalankan program:
   ```
   ./main
   ```

3. Gunakan menu yang tersedia untuk:
   - Menambahkan tim baru
   - Menampilkan daftar tim
   - Mencari tim tertentu
   - Memperbarui data tim
   - Menghapus tim dari sistem
   - Melihat statistik tim
   - Mengurutkan tim berdasarkan kriteria yang dipilih

## Batasan Program

- Program menggunakan array statis dengan kapasitas maksimal 100 tim
- Program tidak menggunakan pernyataan break (selain repeat-until) atau continue
- Program menggunakan pendekatan "soft delete" dimana data yang dihapus ditandai dengan flag `IsDeleted`

## Implementasi Teknis

Proyek ini mengimplementasikan:
- Struktur data array statis
- Tipe bentukan (struct)
- Algoritma sequential search
- Algoritma selection sort dan insertion sort
- Pemrograman modular dengan subprogram (fungsi)
- Fungsi statistik untuk menghitung rata-rata, nilai tertinggi, dan nilai terendah