package main

import (
	"fmt"
	"strings"
)

// Struktur data untuk Tim Esport
type TimEsport struct {
	ID       string
	Nama     string
	Fakultas string
	Poin     int
}

// Variabel global untuk menyimpan data tim
var timData [5]TimEsport // Array statis dengan kapasitas 5 tim
var jumlahTim int = 0    // Jumlah tim yang tersimpan

// Fungsi untuk menambahkan data dummy
func tambahDataDummy() {
	// Pastikan tidak melebihi kapasitas
	if jumlahTim < len(timData) {
		timData[jumlahTim] = TimEsport{ID: "T01", Nama: "Alpha Coders", Fakultas: "Informatika", Poin: 100}
		jumlahTim++
	}
	if jumlahTim < len(timData) {
		timData[jumlahTim] = TimEsport{ID: "T02", Nama: "Beta Testers", Fakultas: "Elektro", Poin: 85}
		jumlahTim++
	}
	if jumlahTim < len(timData) {
		timData[jumlahTim] = TimEsport{ID: "T03", Nama: "Gamma Gamers", Fakultas: "DKV", Poin: 92}
		jumlahTim++
	}
}

// Fungsi utama
func main() {
	tambahDataDummy() // Panggil fungsi untuk menambahkan data dummy di awal

	var pilihan int
	padding := "                                  "

	for {
		tampilkanMenu()
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahTim()
		case 2:
			tampilkanSemuaTim()
		case 3:
			cariTim()
		case 4:
			updateTim()
		case 5:
			hapusTim()
		case 6:
			tampilkanStatistik()
		case 7:
			urutkanTim()
		case 0:
			fmt.Println(padding + "Terima kasih telah menggunakan sistem pendaftaran lomba esport!")
			return
		default:
			fmt.Println(padding + "Pilihan tidak valid. Silakan coba lagi.")
		}
		fmt.Println("\nTekan Enter untuk melanjutkan...")
		fmt.Scanln() // Membersihkan buffer
		var dummy string
		fmt.Scanln(&dummy) // Menunggu Enter
	}
}

// Menampilkan menu utama
func tampilkanMenu() {
	indent := "                                  " // 34 spasi
	fmt.Println(indent + "╔═══════════════════════════════════════════════════════════════════════════╗")
	fmt.Println(indent + "║                       SISTEM PENDAFTARAN LOMBA ESPORT                     ║")
	fmt.Println(indent + "║                Tugas Besar Algortima Pemrograman Smester 2                ║")
	fmt.Println(indent + "║                                                                           ║")
	fmt.Println(indent + "║  Anggota:                                                                 ║")
	fmt.Println(indent + "║   - Harsya Brahmantyo                                                     ║")
	fmt.Println(indent + "║   - Raditya Vihandika Bari Jabran                                         ║")
	fmt.Println(indent + "║   - Muhamad Pradipa Dwi Pahlevi                                           ║")
	fmt.Println(indent + "╠═══════════════════════════════════════════════════════════════════════════╣")
	fmt.Println(indent + "║ 1. Tambah Tim                                                             ║")
	fmt.Println(indent + "║ 2. Lihat Data Tim                                                         ║")
	fmt.Println(indent + "║ 3. Cari Tim                                                               ║")
	fmt.Println(indent + "║ 4. Update Data Tim                                                        ║")
	fmt.Println(indent + "║ 5. Hapus Tim                                                              ║")
	fmt.Println(indent + "║ 6. Statistik Tim                                                          ║")
	fmt.Println(indent + "║ 7. Urutkan Tim                                                            ║")
	fmt.Println(indent + "║ 0. Keluar                                                                 ║")
	fmt.Println(indent + "║                                                             ┌───────────┐ ║")
	fmt.Println(indent + "║                                                             │ MAIN MENU │ ║")
	fmt.Println(indent + "║                                                             └───────────┘ ║")
	fmt.Println(indent + "╚═══════════════════════════════════════════════════════════════════════════╝")
	fmt.Print(indent + "")
}

// Menambahkan tim baru
func tambahTim() {
	if jumlahTim >= len(timData) { // Check against capacity of timData
		indent := "                                   "
		fmt.Println(indent + "Maaf, kapasitas tim sudah penuh!")
		return
	}

	var tim TimEsport
	indent := "                                  "

	fmt.Println(indent + "TAMBAH TIM BARU")

	for {
		fmt.Print(indent + "ID Tim: ")
		fmt.Scan(&tim.ID)

		if cariTimByID(tim.ID) == -1 { // ID must be unique among active teams
			break
		}
		fmt.Println(indent + "ID Tim sudah digunakan. Silakan gunakan ID lain.")
	}

	// Input Nama Tim
	fmt.Print(indent + "Nama Tim: ")
	fmt.Scan(&tim.Nama)
	var namaTambahan string
	fmt.Scanln(&namaTambahan)
	tim.Nama = tim.Nama + namaTambahan
	tim.Nama = strings.TrimSpace(tim.Nama)

	// Input Fakultas
	fmt.Print(indent + "Fakultas: ")
	fmt.Scan(&tim.Fakultas)
	var fakultasTambahan string
	fmt.Scanln(&fakultasTambahan)
	tim.Fakultas = tim.Fakultas + fakultasTambahan
	tim.Fakultas = strings.TrimSpace(tim.Fakultas)

	// Input Poin
	fmt.Print(indent + "Poin: ")
	fmt.Scan(&tim.Poin)

	// Simpan data tim
	timData[jumlahTim] = tim
	jumlahTim++

	fmt.Println(indent + "Tim berhasil ditambahkan!")
}

// Menampilkan semua tim
func tampilkanSemuaTim() {
	if jumlahTim == 0 {
		padding := "                                  "
		fmt.Println(padding + "Belum ada tim yang terdaftar...")
		return
	}

	padding := "                                  "

	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════════════════════════════════════════╗")
	fmt.Println(padding + "║ ID  | Nama Tim             | Fakultas             | Poin          ║")
	fmt.Println(padding + "╠═══════════════════════════════════════════════════════════════════╣")

	for i := 0; i < jumlahTim; i++ {
		fmt.Printf(padding+"║ %-3s | %-20s | %-20s | %-13d ║\n",
			timData[i].ID,
			timData[i].Nama,
			timData[i].Fakultas,
			timData[i].Poin)
	}
	fmt.Println(padding + "╚═══════════════════════════════════════════════════════════════════╝")
}

// Mencari tim berdasarkan ID (sequential search)
func cariTimByID(id string) int {
	for i := 0; i < jumlahTim; i++ {
		if timData[i].ID == id {
			return i
		}
	}
	return -1
}

// Mencari tim berdasarkan nama (sequential search)
func cariTimByNama(nama string) int {
	for i := 0; i < jumlahTim; i++ {
		if strings.ToLower(timData[i].Nama) == strings.ToLower(nama) {
			return i
		}
	}
	return -1
}

// Binary search by ID (assumes slice is sorted by ID)
func binarySearchByID(slice []TimEsport, id string) int {
	low := 0
	high := len(slice) - 1
	for low <= high {
		mid := low + (high-low)/2
		if slice[mid].ID == id {
			return mid
		}
		if slice[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // Not found
}

// Binary search by Name (assumes slice is sorted by Name, case-insensitive)
func binarySearchByNama(slice []TimEsport, nama string) int {
	targetNama := strings.ToLower(nama)
	low := 0
	high := len(slice) - 1
	for low <= high {
		mid := low + (high-low)/2
		currentNama := strings.ToLower(slice[mid].Nama)
		if currentNama == targetNama {
			return mid
		}
		if currentNama < targetNama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // Not found
}

// Mencari tim (menu)
func cariTim() {
	var searchFieldChoice int
	padding := "                                  "

	if jumlahTim == 0 {
		fmt.Println(padding + "Belum ada tim yang terdaftar untuk dicari.")
		return
	}

	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════╗")
	fmt.Println(padding + "║           CARI TIM            ║")
	fmt.Println(padding + "╠═══════════════════════════════╣")
	fmt.Println(padding + "║ 1. Cari berdasarkan ID Tim    ║")
	fmt.Println(padding + "║ 2. Cari berdasarkan Nama Tim  ║")
	fmt.Println(padding + "╚═══════════════════════════════╝")
	fmt.Print(padding + "Pilihan Anda: ")
	fmt.Scan(&searchFieldChoice)

	var algorithmChoice int
	fmt.Println()
	fmt.Println(padding + "╔════════════════════════════════╗")
	fmt.Println(padding + "║      PILIH ALGORITMA CARI      ║")
	fmt.Println(padding + "╠════════════════════════════════╣")
	fmt.Println(padding + "║ 1. Sequential Search           ║")
	fmt.Println(padding + "║ 2. Binary Search               ║")
	fmt.Println(padding + "╚════════════════════════════════╝")
	fmt.Print(padding + "Pilihan Algoritma: ")
	fmt.Scan(&algorithmChoice)

	switch searchFieldChoice {
	case 1: // Search by ID
		var id string
		fmt.Print(padding + "Masukkan ID Tim: ")
		fmt.Scan(&id)

		if algorithmChoice == 1 { // Sequential Search by ID
			idx := cariTimByID(id)
			if idx != -1 {
				tampilkanDetailTimByIndex(idx)
			} else {
				fmt.Println(padding + "Tim dengan ID tersebut tidak ditemukan (Sequential).")
			}
		} else if algorithmChoice == 2 { // Binary Search by ID
			searchableTeams := make([]TimEsport, jumlahTim)
			copy(searchableTeams, timData[:jumlahTim])
			selectionSortByID(searchableTeams, true) // Sort ascending for binary search

			idxInSortedSlice := binarySearchByID(searchableTeams, id)
			if idxInSortedSlice != -1 {
				tampilkanDetailTimData(searchableTeams[idxInSortedSlice])
			} else {
				fmt.Println(padding + "Tim dengan ID tersebut tidak ditemukan (Binary).")
			}
		} else {
			fmt.Println(padding + "Pilihan algoritma tidak valid.")
		}

	case 2: // Search by Name
		var nama string
		fmt.Print(padding + "Masukkan Nama Tim: ")
		fmt.Scan(&nama)
		var namaTambahan string
		fmt.Scanln(&namaTambahan)
		nama = nama + namaTambahan
		nama = strings.TrimSpace(nama)

		if algorithmChoice == 1 { // Sequential Search by Name
			idx := cariTimByNama(nama)
			if idx != -1 {
				tampilkanDetailTimByIndex(idx)
			} else {
				fmt.Println(padding + "Tim dengan nama tersebut tidak ditemukan (Sequential).")
			}
		} else if algorithmChoice == 2 { // Binary Search by Name
			searchableTeams := make([]TimEsport, jumlahTim)
			copy(searchableTeams, timData[:jumlahTim])
			selectionSortByNama(searchableTeams, true) // Sort ascending for binary search

			idxInSortedSlice := binarySearchByNama(searchableTeams, nama)
			if idxInSortedSlice != -1 {
				tampilkanDetailTimData(searchableTeams[idxInSortedSlice])
			} else {
				fmt.Println(padding + "Tim dengan nama tersebut tidak ditemukan (Binary).")
			}
		} else {
			fmt.Println(padding + "Pilihan algoritma tidak valid.")
		}
	default:
		fmt.Println(padding + "Pilihan field pencarian tidak valid.")
	}
}

// Menampilkan detail tim berdasarkan indeks di timData
func tampilkanDetailTimByIndex(idx int) {
	padding := "                                  "
	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════╗")
	fmt.Println(padding + "║           DETAIL TIM          ║")
	fmt.Println(padding + "╠═══════════════════════════════╣")
	fmt.Printf(padding+"║ ID Tim    : %-17s ║\n", timData[idx].ID)
	fmt.Printf(padding+"║ Nama Tim  : %-17s ║\n", timData[idx].Nama)
	fmt.Printf(padding+"║ Fakultas  : %-17s ║\n", timData[idx].Fakultas)
	fmt.Printf(padding+"║ Poin      : %-17d ║\n", timData[idx].Poin)
	fmt.Println(padding + "╚═══════════════════════════════╝")
}

// Menampilkan detail tim dari struct TimEsport (untuk hasil binary search)
func tampilkanDetailTimData(tim TimEsport) {
	padding := "                                  "
	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════╗")
	fmt.Println(padding + "║           DETAIL TIM          ║")
	fmt.Println(padding + "╠═══════════════════════════════╣")
	fmt.Printf(padding+"║ ID Tim    : %-17s ║\n", tim.ID)
	fmt.Printf(padding+"║ Nama Tim  : %-17s ║\n", tim.Nama)
	fmt.Printf(padding+"║ Fakultas  : %-17s ║\n", tim.Fakultas)
	fmt.Printf(padding+"║ Poin      : %-17d ║\n", tim.Poin)
	fmt.Println(padding + "╚═══════════════════════════════╝")
}

// Update data tim
func updateTim() {
	var id string
	padding := "                                  "
	fmt.Println(padding + "UPDATE DATA TIM")

	if jumlahTim == 0 {
		fmt.Println(padding + "Belum ada tim yang terdaftar untuk diupdate.")
		return
	}

	fmt.Print(padding + "Masukkan ID Tim yang akan diupdate: ")
	fmt.Scan(&id)

	idx := cariTimByID(id)
	if idx == -1 {
		fmt.Println("\nTim dengan ID tersebut tidak ditemukan.")
		return
	}

	tampilkanDetailTimByIndex(idx)

	var pilihan int
	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════╗")
	fmt.Println(padding + "║  Data apa yang ingin diubah?  ║")
	fmt.Println(padding + "╠═══════════════════════════════╣")
	fmt.Println(padding + "║ 1. Nama Tim                   ║")
	fmt.Println(padding + "║ 2. Fakultas                   ║")
	fmt.Println(padding + "║ 3. Poin                       ║")
	fmt.Println(padding + "╚═══════════════════════════════╝")
	fmt.Print(padding + "Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var nama string
		fmt.Print(padding + "Nama Tim Baru: ")
		fmt.Scan(&nama)
		var namaTambahan string
		fmt.Scanln(&namaTambahan)
		nama = nama + namaTambahan
		timData[idx].Nama = strings.TrimSpace(nama)
		fmt.Println(padding + "Nama tim berhasil diupdate!")
	case 2:
		var fakultas string
		fmt.Print(padding + "Fakultas Baru: ")
		fmt.Scan(&fakultas)
		var fakultasTambahan string
		fmt.Scanln(&fakultasTambahan)
		fakultas = fakultas + fakultasTambahan
		timData[idx].Fakultas = strings.TrimSpace(fakultas)
		fmt.Println(padding + "Fakultas berhasil diupdate!")
	case 3:
		var poin int
		fmt.Print(padding + "Poin Baru: ")
		fmt.Scan(&poin)
		timData[idx].Poin = poin
		fmt.Println(padding + "Poin berhasil diupdate!")
	default:
		fmt.Println(padding + "Pilihan tidak valid.")
	}
}

// Hapus tim
func hapusTim() {
	var id string
	padding := "                                  "

	fmt.Println(padding + "HAPUS TIM")

	if jumlahTim == 0 {
		fmt.Println(padding + "Belum ada tim yang terdaftar untuk dihapus.")
		return
	}

	fmt.Print(padding + "Masukkan ID Tim yang akan dihapus: ")
	fmt.Scan(&id)

	idx := cariTimByID(id)
	if idx == -1 {
		fmt.Println(padding + "Tim dengan ID tersebut tidak ditemukan.")
		return
	}

	tampilkanDetailTimByIndex(idx)

	var konfirmasi string
	fmt.Print(padding + "Apakah Anda yakin ingin menghapus tim ini? (y/n): ")
	fmt.Scan(&konfirmasi)

	if strings.ToLower(konfirmasi) == "y" {
		// Shift elements to the left to fill the gap
		for i := idx; i < jumlahTim-1; i++ {
			timData[i] = timData[i+1]
		}
		jumlahTim--                      // Decrease the total count of teams
		timData[jumlahTim] = TimEsport{} // Clear the last element (now outside active range)
		fmt.Println(padding + "Tim berhasil dihapus!")
	} else {
		fmt.Println(padding + "Penghapusan dibatalkan.")
	}
}

// Menampilkan statistik tim
func tampilkanStatistik() {
	padding := "                                  "
	if jumlahTim == 0 {
		fmt.Println(padding + "Belum ada tim yang terdaftar.")
		return
	}

	var totalPoin int
	minPoin := timData[0].Poin
	maxPoin := timData[0].Poin
	timMinPoin := timData[0].Nama
	timMaxPoin := timData[0].Nama

	totalPoin += timData[0].Poin

	for i := 1; i < jumlahTim; i++ { // Start from 1 as 0 is initialized
		if timData[i].Poin < minPoin {
			minPoin = timData[i].Poin
			timMinPoin = timData[i].Nama
		}
		if timData[i].Poin > maxPoin {
			maxPoin = timData[i].Poin
			timMaxPoin = timData[i].Nama
		}
		totalPoin += timData[i].Poin
	}

	rataRataPoin := float64(totalPoin) / float64(jumlahTim)

	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════╗")
	fmt.Println(padding + "║         STATISTIK TIM         ║")
	fmt.Println(padding + "╚═══════════════════════════════╝")
	fmt.Println(padding+"Jumlah Tim Terdaftar :", jumlahTim)
	fmt.Printf(padding+"Rata-rata Poin       : %.2f\n", rataRataPoin)
	fmt.Println(padding+"Poin Tertinggi       :", maxPoin, "(", timMaxPoin, ")")
	fmt.Println(padding+"Poin Terendah        :", minPoin, "(", timMinPoin, ")")
}

// Urutkan tim
func urutkanTim() {
	var pilihan, jenisUrutan int
	padding := "                                  "

	if jumlahTim == 0 {
		fmt.Println(padding + "Belum ada tim yang terdaftar untuk diurutkan.")
		return
	}

	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════╗")
	fmt.Println(padding + "║        URUTKAN TIM            ║")
	fmt.Println(padding + "╠═══════════════════════════════╣")
	fmt.Println(padding + "║ 1. ID Tim                     ║")
	fmt.Println(padding + "║ 2. Nama Tim                   ║")
	fmt.Println(padding + "║ 3. Fakultas                   ║")
	fmt.Println(padding + "║ 4. Poin                       ║")
	fmt.Println(padding + "╚═══════════════════════════════╝")
	fmt.Print(padding + "Pilihan Anda: ")
	fmt.Scan(&pilihan)

	fmt.Println()
	fmt.Println(padding + "╔════════════════════════════════════════════╗")
	fmt.Println(padding + "║             JENIS PENGURUTAN               ║")
	fmt.Println(padding + "╠════════════════════════════════════════════╣")
	fmt.Println(padding + "║ 1. Ascending (A-Z / 0-9)                   ║")
	fmt.Println(padding + "║ 2. Descending (Z-A / 9-0)                  ║")
	fmt.Println(padding + "╚════════════════════════════════════════════╝")
	fmt.Print(padding + "Pilihan Anda: ")
	fmt.Scan(&jenisUrutan)

	isAscending := (jenisUrutan == 1)
	var algoritma int

	fmt.Println()
	fmt.Println(padding + "╔════════════════════════════════╗")
	fmt.Println(padding + "║       ALGORITMA PENGURUTAN     ║")
	fmt.Println(padding + "╠════════════════════════════════╣")
	fmt.Println(padding + "║ 1. Selection Sort              ║")
	fmt.Println(padding + "║ 2. Insertion Sort              ║")
	fmt.Println(padding + "╚════════════════════════════════╝")
	fmt.Print(padding + "Pilihan Anda: ")
	fmt.Scan(&algoritma)

	// Create a copy of active teams to sort
	timToSort := make([]TimEsport, jumlahTim)
	copy(timToSort, timData[:jumlahTim]) // Copy from global data up to jumlahTim

	if len(timToSort) == 0 {
		fmt.Println(padding + "Belum ada tim aktif yang terdaftar.")
		return
	}

	switch algoritma {
	case 1: // Selection Sort
		switch pilihan {
		case 1:
			selectionSortByID(timToSort, isAscending)
		case 2:
			selectionSortByNama(timToSort, isAscending)
		case 3:
			selectionSortByFakultas(timToSort, isAscending)
		case 4:
			selectionSortByPoin(timToSort, isAscending)
		default:
			fmt.Println(padding + "Pilihan field pengurutan tidak valid.")
			return
		}
	case 2: // Insertion Sort
		switch pilihan {
		case 1:
			insertionSortByID(timToSort, isAscending)
		case 2:
			insertionSortByNama(timToSort, isAscending)
		case 3:
			insertionSortByFakultas(timToSort, isAscending)
		case 4:
			insertionSortByPoin(timToSort, isAscending)
		default:
			fmt.Println(padding + "Pilihan field pengurutan tidak valid.")
			return
		}
	default:
		fmt.Println(padding + "Pilihan algoritma pengurutan tidak valid.")
		return
	}

	// Tampilkan hasil pengurutan
	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════════════════════════════════════════╗")
	fmt.Println(padding + "║ ID  | Nama Tim             | Fakultas             | Poin          ║")
	fmt.Println(padding + "╠═══════════════════════════════════════════════════════════════════╣")

	for i := 0; i < len(timToSort); i++ {
		fmt.Printf(padding+"║ %-3s | %-20s | %-20s | %-13d ║\n",
			timToSort[i].ID,
			timToSort[i].Nama,
			timToSort[i].Fakultas,
			timToSort[i].Poin)
	}
	fmt.Println(padding + "╚═══════════════════════════════════════════════════════════════════╝")
}

// Implementasi Selection Sort

func selectionSortByID(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		targetIdx := i
		for j := i + 1; j < n; j++ {
			if isAscending {
				if arr[j].ID < arr[targetIdx].ID {
					targetIdx = j
				}
			} else {
				if arr[j].ID > arr[targetIdx].ID {
					targetIdx = j
				}
			}
		}
		if targetIdx != i {
			arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
		}
	}
}

func selectionSortByNama(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		targetIdx := i
		for j := i + 1; j < n; j++ {
			if isAscending {
				if strings.ToLower(arr[j].Nama) < strings.ToLower(arr[targetIdx].Nama) {
					targetIdx = j
				}
			} else {
				if strings.ToLower(arr[j].Nama) > strings.ToLower(arr[targetIdx].Nama) {
					targetIdx = j
				}
			}
		}
		if targetIdx != i {
			arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
		}
	}
}

func selectionSortByFakultas(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		targetIdx := i
		for j := i + 1; j < n; j++ {
			if isAscending {
				if strings.ToLower(arr[j].Fakultas) < strings.ToLower(arr[targetIdx].Fakultas) {
					targetIdx = j
				}
			} else {
				if strings.ToLower(arr[j].Fakultas) > strings.ToLower(arr[targetIdx].Fakultas) {
					targetIdx = j
				}
			}
		}
		if targetIdx != i {
			arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
		}
	}
}

func selectionSortByPoin(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		targetIdx := i
		for j := i + 1; j < n; j++ {
			if isAscending {
				if arr[j].Poin < arr[targetIdx].Poin {
					targetIdx = j
				}
			} else {
				if arr[j].Poin > arr[targetIdx].Poin {
					targetIdx = j
				}
			}
		}
		if targetIdx != i {
			arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
		}
	}
}

// Implementasi Insertion Sort

func insertionSortByID(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		if isAscending {
			for j >= 0 && arr[j].ID > key.ID {
				arr[j+1] = arr[j]
				j--
			}
		} else {
			for j >= 0 && arr[j].ID < key.ID {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
	}
}

func insertionSortByNama(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		if isAscending {
			for j >= 0 && strings.ToLower(arr[j].Nama) > strings.ToLower(key.Nama) {
				arr[j+1] = arr[j]
				j--
			}
		} else {
			for j >= 0 && strings.ToLower(arr[j].Nama) < strings.ToLower(key.Nama) {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
	}
}

func insertionSortByFakultas(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		if isAscending {
			for j >= 0 && strings.ToLower(arr[j].Fakultas) > strings.ToLower(key.Fakultas) {
				arr[j+1] = arr[j]
				j--
			}
		} else {
			for j >= 0 && strings.ToLower(arr[j].Fakultas) < strings.ToLower(key.Fakultas) {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
	}
}

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
