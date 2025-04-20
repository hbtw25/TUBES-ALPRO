package main

import (
	"fmt"
)

// Initial State (IS):
// --------------------
// - jumlahTim = 0
// - Array timData kosong
// - Belum ada tim yang ditambahkan
//
// Final State (FS):
// ------------------
// - Setelah pengguna menambahkan tim:
//   - jumlahTim > 0
//   - timData berisi data tim yang ditambahkan
//   - Program dapat menampilkan dan mencari tim berdasarkan ID
//

// Struct untuk menyimpan data tim esport
type TimEsport struct {
	ID        string // ID unik tim
	Nama      string // Nama tim
	Fakultas  string // Asal fakultas
	Poin      int    // Jumlah poin
	IsDeleted bool   // Penanda soft delete
}

// Array statis untuk menyimpan maksimal 100 tim
var timData [100]TimEsport
var jumlahTim int = 0 // IS: jumlahTim = 0 (belum ada tim)

func main() {
	var pilihan int

	for {
		tampilkanMenu()
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahTim() // Tambah data tim ke array
		case 2:
			tampilkanSemuaTim() // Menampilkan seluruh data tim
		case 3:
			cariTim() // Mencari tim berdasarkan ID
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}

		// Jeda sebelum kembali ke menu
		fmt.Println("\nTekan Enter untuk melanjutkan...")
		fmt.Scanln()
		var dummy string
		fmt.Scanln(&dummy)
	}
}

// Menampilkan menu utama program
func tampilkanMenu() {
	fmt.Println("\n========================================")
	fmt.Println("   SISTEM PENDAFTARAN LOMBA ESPORT")
	fmt.Println("========================================")
	fmt.Println("1. Tambah Tim Baru")
	fmt.Println("2. Tampilkan Semua Tim")
	fmt.Println("3. Cari Tim")
	fmt.Println("0. Keluar")
	fmt.Println("========================================")
}

// Tambahkan tim ke array timData
func tambahTim() {
	if jumlahTim >= 100 {
		fmt.Println("Kapasitas penuh!")
		return
	}

	var tim TimEsport
	fmt.Print("ID Tim: ")
	fmt.Scan(&tim.ID)
	fmt.Print("Nama Tim: ")
	fmt.Scan(&tim.Nama)
	fmt.Print("Fakultas: ")
	fmt.Scan(&tim.Fakultas)
	fmt.Print("Poin: ")
	fmt.Scan(&tim.Poin)

	timData[jumlahTim] = tim
	jumlahTim++ // FS: jumlahTim bertambah setelah tambah tim

	fmt.Println("Tim berhasil ditambahkan!")
}

// Menampilkan seluruh data tim yang belum dihapus
func tampilkanSemuaTim() {
	if jumlahTim == 0 {
		fmt.Println("Belum ada tim.")
		return
	}

	fmt.Println("\nDaftar Tim:")
	for i := 0; i < jumlahTim; i++ {
		if !timData[i].IsDeleted {
			fmt.Printf("%s - %s (%s) - %d poin\n", timData[i].ID, timData[i].Nama, timData[i].Fakultas, timData[i].Poin)
		}
	}
}

// Cari tim berdasarkan ID
func cariTim() {
	var id string
	fmt.Print("Masukkan ID Tim: ")
	fmt.Scan(&id)

	for i := 0; i < jumlahTim; i++ {
		if timData[i].ID == id && !timData[i].IsDeleted {
			fmt.Printf("Ditemukan: %s - %s (%s) - %d poin\n", timData[i].ID, timData[i].Nama, timData[i].Fakultas, timData[i].Poin)
			return
		}
	}
	fmt.Println("Tim tidak ditemukan.")
}
