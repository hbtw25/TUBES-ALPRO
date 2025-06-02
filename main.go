package main

import (
	"fmt"
	"strings"
)

// Struktur data untuk Tim Esport
type TimEsport struct {
	ID        string
	Nama      string
	Fakultas  string
	Poin      int
	IsDeleted bool // Flag untuk penanda data yang sudah dihapus
}

// Variabel global untuk menyimpan data tim
var timData [5]TimEsport // Array statis dengan kapasitas 100 tim
var jumlahTim int = 0      // Jumlah tim yang tersimpan

// Fungsi utama
func main() {
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
			fmt.Println(paddding + "Terima kasih telah menggunakan sistem pendaftaran lomba esport!")
			return
		default:
			fmt.Println(padding + "Pilihan tidak valid. Silakan coba lagi.")
		}
		fmt.Println("\nTekan Enter untuk melanjutkan...")
		fmt.Scanln() // Membersihkan buffer
		var dummy string
		fmt.Scanln(&dummy)
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
	if jumlahTim >= 5 {
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
		
		if cariTimByID(tim.ID) == -1 {
			break
		}
		fmt.Println(indent + "ID Tim sudah digunakan. Silakan gunakan ID lain.")
	}
	
	// Input Nama Tim
	fmt.Print(indent + "Nama Tim: ")
	fmt.Scan(&tim.Nama)
	var namaTambahan string
	fmt.Scanln(&namaTambahan) // Menangkap sisa input setelah spasi
	tim.Nama = tim.Nama + namaTambahan
	tim.Nama = strings.TrimSpace(tim.Nama)
	
	// Input Fakultas
	fmt.Print(indent + "Fakultas: ")
	fmt.Scan(&tim.Fakultas)
	var fakultasTambahan string
	fmt.Scanln(&fakultasTambahan) // Menangkap sisa input setelah spasi
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
	fmt.Println(padding + "║ ID  |      Nama Tim        |      Fakultas        | Poin          ║")
	fmt.Println(padding + "╠═══════════════════════════════════════════════════════════════════╣")

	for i := 0; i < jumlahTim; i++ {
		if !timData[i].IsDeleted {
			fmt.Printf(padding+"║ %-3s | %-20s | %-20s | %-13d ║\n",
				timData[i].ID,
				timData[i].Nama,
				timData[i].Fakultas,
				timData[i].Poin)
		}
	}
	fmt.Println(padding + "╚═══════════════════════════════════════════════════════════════════╝")
}

// Mencari tim berdasarkan ID (sequential search)
func cariTimByID(id string) int {
	for i := 0; i < jumlahTim; i++ {
		if timData[i].ID == id && !timData[i].IsDeleted {
			return i
		}
	}
	return -1
}

// Mencari tim berdasarkan nama (sequential search)
func cariTimByNama(nama string) int {
	for i := 0; i < jumlahTim; i++ {
		if strings.ToLower(timData[i].Nama) == strings.ToLower(nama) && !timData[i].IsDeleted {
			return i
		}
	}
	return -1
}

// Mencari tim (menu)
func cariTim() {
	var pilihan int
	
	padding := "                                  "
	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════╗")
	fmt.Println(padding + "║           CARI TIM            ║")
	fmt.Println(padding + "╠═══════════════════════════════╣")
	fmt.Println(padding + "║ 1. Cari berdasarkan ID Tim    ║")
	fmt.Println(padding + "║ 2. Cari berdasarkan Nama Tim  ║")
	fmt.Println(padding + "╚═══════════════════════════════╝")
	fmt.Print(padding + "Pilihan Anda: ")
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case 1:
		var id string
		fmt.Print(padding + "Masukkan ID Tim: ")
		fmt.Scan(&id)
		
		idx := cariTimByID(id)
		if idx != -1 {
			tampilkanDetailTim(idx)
		} else {
			fmt.Println(padding + "Tim dengan ID tersebut tidak ditemukan.")
		}
	case 2:
		var nama string
		fmt.Print(padding + "Masukkan Nama Tim: ")
		fmt.Scan(&nama)
		var namaTambahan string
		fmt.Scanln(&namaTambahan)
		nama = nama + namaTambahan
		nama = strings.TrimSpace(nama)
		
		idx := cariTimByNama(nama)
		if idx != -1 {
			tampilkanDetailTim(idx)
		} else {
			fmt.Println(padding + "Tim dengan nama tersebut tidak ditemukan.")
		}
	default:
		fmt.Println(padding + "Pilihan tidak valid.")
	}
}

// Menampilkan detail tim berdasarkan indeks
func tampilkanDetailTim(idx int) {
	padding := "                                  " // 34 spasi, sesuaikan jika perlu
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

// Update data tim
func updateTim() {
	var id string
	padding := "                                  "
	fmt.Println(padding + "UPDATE DATA TIM")
	fmt.Print(padding + "Masukkan ID Tim yang akan diupdate: ")
	fmt.Scan(&id)
	
	idx := cariTimByID(id)
	if idx == -1 {
		fmt.Println("\nTim dengan ID tersebut tidak ditemukan.")
		return
	}
	
	tampilkanDetailTim(idx)
	
	var pilihan int
	padding = "                                  "
	
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
	fmt.Print(padding + "Masukkan ID Tim yang akan dihapus: ")
	fmt.Scan(&id)
	
	idx := cariTimByID(id)
	if idx == -1 {
		fmt.Println(padding + "Tim dengan ID tersebut tidak ditemukan.")
		return
	}
	
	tampilkanDetailTim(idx)
	
	var konfirmasi string
	fmt.Print(padding + "Apakah Anda yakin ingin menghapus tim ini? (y/n): ")
	fmt.Scan(&konfirmasi)
	
	if strings.ToLower(konfirmasi) == "y" {
		timData[idx].IsDeleted = true
		fmt.Println(padding + "Tim berhasil dihapus!")
	} else {
		fmt.Println(padding + "Penghapusan dibatalkan.")
	}
}

// Menampilkan statistik tim
func tampilkanStatistik() {
	if jumlahTim == 0 {
		padding := "                                  "
		fmt.Println(padding + "Belum ada tim yang terdaftar.")
		return
	}
	
	var totalPoin, jumlahTimAktif int
	var minPoin, maxPoin int
	var timMinPoin, timMaxPoin string
	
	// Set nilai awal
	var adaTimAktif bool = false
	
	for i := 0; i < jumlahTim; i++ {
		if !timData[i].IsDeleted {
			if !adaTimAktif {
				// Inisialisasi nilai awal
				minPoin = timData[i].Poin
				maxPoin = timData[i].Poin
				timMinPoin = timData[i].Nama
				timMaxPoin = timData[i].Nama
				adaTimAktif = true
			} else {
				if timData[i].Poin < minPoin {
					minPoin = timData[i].Poin
					timMinPoin = timData[i].Nama
				}
				if timData[i].Poin > maxPoin {
					maxPoin = timData[i].Poin
					timMaxPoin = timData[i].Nama
				}
			}
			
			totalPoin += timData[i].Poin
			jumlahTimAktif++
		}
	}
	
	if jumlahTimAktif == 0 {
		padding := "                                  "
		
		fmt.Println(padding + "Belum ada tim aktif yang terdaftar.")
		return
	}
	
	rataRataPoin := float64(totalPoin) / float64(jumlahTimAktif)
	
	padding := "                                  "
	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════╗")
	fmt.Println(padding + "║         STATISTIK TIM         ║")
	fmt.Println(padding + "╚═══════════════════════════════╝")
	fmt.Println(padding + "Jumlah Tim Terdaftar :", jumlahTimAktif)
	fmt.Printf(padding + "Rata-rata Poin       : %.2f\n", rataRataPoin)
	fmt.Println(padding + "Poin Tertinggi       :", maxPoin, "(", timMaxPoin, ")")
	fmt.Println(padding + "Poin Terendah        :", minPoin, "(", timMinPoin, ")")
}

// Urutkan tim
func urutkanTim() {
	var pilihan, jenisUrutan int
	padding := "                                  " // 34 spasi, sesuaikan jika perlu
	
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
	
	var isAscending bool = (jenisUrutan == 1)
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
	
	var timAktif []TimEsport
	for i := 0; i < jumlahTim; i++ {
		if !timData[i].IsDeleted {
			timAktif = append(timAktif, timData[i])
		}
	}
	
	if len(timAktif) == 0 {
		fmt.Println(padding + "Belum ada tim aktif yang terdaftar.")
		return
	}
	
	switch algoritma {
	case 1:
		switch pilihan {
		case 1:
			selectionSortByID(timAktif, isAscending)
		case 2:
			selectionSortByNama(timAktif, isAscending)
		case 3:
			selectionSortByFakultas(timAktif, isAscending)
		case 4:
			selectionSortByPoin(timAktif, isAscending)
		default:
			fmt.Println(padding + "Pilihan tidak valid.")
			return
		}
	case 2:
		switch pilihan {
		case 1:
			insertionSortByID(timAktif, isAscending)
		case 2:
			insertionSortByNama(timAktif, isAscending)
		case 3:
			insertionSortByFakultas(timAktif, isAscending)
		case 4:
			insertionSortByPoin(timAktif, isAscending)
		default:
			fmt.Println(padding + "Pilihan tidak valid.")
			return
		}
	default:
		fmt.Println(padding + "Pilihan tidak valid.")
		return
	}
	
	// Tampilkan hasil pengurutan
	fmt.Println()
	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════════════════════════════════════════╗")
	fmt.Println(padding + "║ ID  |      Nama Tim        |      Fakultas        | Poin          ║")
	fmt.Println(padding + "╠═══════════════════════════════════════════════════════════════════╣")

	for i := 0; i < len(timAktif); i++ {
		fmt.Printf(padding+"║ %-3s | %-20s | %-20s | %-13d ║\n",
			timAktif[i].ID,
			timAktif[i].Nama,
			timAktif[i].Fakultas,
			timAktif[i].Poin)
	}
	fmt.Println(padding + "╚═══════════════════════════════════════════════════════════════════╝")
}


// Implementasi Selection Sort

func selectionSortByID(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if isAscending {
				if arr[j].ID < arr[minIdx].ID {
					minIdx = j
				}
			} else {
				if arr[j].ID > arr[minIdx].ID {
					minIdx = j
				}
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

func selectionSortByNama(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if isAscending {
				if strings.ToLower(arr[j].Nama) < strings.ToLower(arr[minIdx].Nama) {
					minIdx = j
				}
			} else {
				if strings.ToLower(arr[j].Nama) > strings.ToLower(arr[minIdx].Nama) {
					minIdx = j
				}
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

func selectionSortByFakultas(arr []TimEsport, isAscending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if isAscending {
				if strings.ToLower(arr[j].Fakultas) < strings.ToLower(arr[minIdx].Fakultas) {
					minIdx = j
				}
			} else {
				if strings.ToLower(arr[j].Fakultas) > strings.ToLower(arr[minIdx].Fakultas) {
					minIdx = j
				}
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

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
