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
var timData [100]TimEsport // Array statis dengan kapasitas 100 tim
var jumlahTim int = 0      // Jumlah tim yang tersimpan

// Fungsi utama
func main() {
	var pilihan int

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
			fmt.Println("\nTerima kasih telah menggunakan sistem pendaftaran lomba esport!")
			return
		default:
			fmt.Println("\nPilihan tidak valid. Silakan coba lagi.")
		}
		
		fmt.Println("\nTekan Enter untuk melanjutkan...")
		fmt.Scanln() // Membersihkan buffer
		var dummy string
		fmt.Scanln(&dummy)
	}
}

// Menampilkan menu utama
func tampilkanMenu() {
	fmt.Println("\n========================================")
	fmt.Println("   SISTEM PENDAFTARAN LOMBA ESPORT")
	fmt.Println("========================================")
	fmt.Println("1. Tambah Tim Baru")
	fmt.Println("2. Tampilkan Semua Tim")
	fmt.Println("3. Cari Tim")
	fmt.Println("4. Update Data Tim")
	fmt.Println("5. Hapus Tim")
	fmt.Println("6. Tampilkan Statistik Tim")
	fmt.Println("7. Urutkan Tim")
	fmt.Println("0. Keluar")
	fmt.Println("========================================")
}

// Menambahkan tim baru
func tambahTim() {
	if jumlahTim >= 100 {
		fmt.Println("\nMaaf, kapasitas tim sudah penuh!")
		return
	}
	
	var tim TimEsport //variabel tim esport
	
	fmt.Println("\n--- TAMBAH TIM BARU ---")
	
	// Input ID Tim
	for {
		fmt.Print("ID Tim: ")
		fmt.Scan(&tim.ID)
		
		if cariTimByID(tim.ID) == -1 {
			break
		}
		fmt.Println("ID Tim sudah digunakan. Silakan gunakan ID lain.")
	}
	
	// Input Nama Tim
	fmt.Print("Nama Tim: ")
	fmt.Scan(&tim.Nama)
	var namaTambahan string
	fmt.Scanln(&namaTambahan) // Menangkap sisa input setelah spasi
	tim.Nama = tim.Nama + namaTambahan
	tim.Nama = strings.TrimSpace(tim.Nama)
	
	// Input Fakultas
	fmt.Print("Fakultas: ")
	fmt.Scan(&tim.Fakultas)
	var fakultasTambahan string
	fmt.Scanln(&fakultasTambahan) // Menangkap sisa input setelah spasi
	tim.Fakultas = tim.Fakultas + fakultasTambahan
	tim.Fakultas = strings.TrimSpace(tim.Fakultas)
	
	// Input Poin
	fmt.Print("Poin: ")
	fmt.Scan(&tim.Poin)
	
	// Simpan data tim
	timData[jumlahTim] = tim
	jumlahTim++
	
	fmt.Println("\nTim berhasil ditambahkan!")
}

// Menampilkan semua tim
func tampilkanSemuaTim() {
	if jumlahTim == 0 {
		fmt.Println("\nBelum ada tim yang terdaftar.")
		return
	}
	
	fmt.Println("\n--- DAFTAR SEMUA TIM ---")
	fmt.Printf("%-5s %-20s %-20s %-5s\n", "ID", "Nama Tim", "Fakultas", "Poin")
	fmt.Println("---------------------------------------------------")
	
	for i := 0; i < jumlahTim; i++ {
		if !timData[i].IsDeleted {
			fmt.Printf("%-5s %-20s %-20s %-5d\n", 
				timData[i].ID, 
				timData[i].Nama, 
				timData[i].Fakultas, 
				timData[i].Poin)
		}
	}
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
	
	fmt.Println("\n--- CARI TIM ---")
	fmt.Println("1. Cari berdasarkan ID Tim")
	fmt.Println("2. Cari berdasarkan Nama Tim")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case 1:
		var id string
		fmt.Print("Masukkan ID Tim: ")
		fmt.Scan(&id)
		
		idx := cariTimByID(id)
		if idx != -1 {
			tampilkanDetailTim(idx)
		} else {
			fmt.Println("\nTim dengan ID tersebut tidak ditemukan.")
		}
	case 2:
		var nama string
		fmt.Print("Masukkan Nama Tim: ")
		fmt.Scan(&nama)
		var namaTambahan string
		fmt.Scanln(&namaTambahan)
		nama = nama + namaTambahan
		nama = strings.TrimSpace(nama)
		
		idx := cariTimByNama(nama)
		if idx != -1 {
			tampilkanDetailTim(idx)
		} else {
			fmt.Println("\nTim dengan nama tersebut tidak ditemukan.")
		}
	default:
		fmt.Println("\nPilihan tidak valid.")
	}
}

// Menampilkan detail tim berdasarkan indeks
func tampilkanDetailTim(idx int) {
	fmt.Println("\n--- DETAIL TIM ---")
	fmt.Println("ID Tim    :", timData[idx].ID)
	fmt.Println("Nama Tim  :", timData[idx].Nama)
	fmt.Println("Fakultas  :", timData[idx].Fakultas)
	fmt.Println("Poin      :", timData[idx].Poin)
}

// Update data tim
func updateTim() {
	var id string
	
	fmt.Println("\n--- UPDATE DATA TIM ---")
	fmt.Print("Masukkan ID Tim yang akan diupdate: ")
	fmt.Scan(&id)
	
	idx := cariTimByID(id)
	if idx == -1 {
		fmt.Println("\nTim dengan ID tersebut tidak ditemukan.")
		return
	}
	
	tampilkanDetailTim(idx)
	
	var pilihan int
	fmt.Println("\nData apa yang ingin diupdate?")
	fmt.Println("1. Nama Tim")
	fmt.Println("2. Fakultas")
	fmt.Println("3. Poin")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case 1:
		var nama string
		fmt.Print("Nama Tim Baru: ")
		fmt.Scan(&nama)
		var namaTambahan string
		fmt.Scanln(&namaTambahan)
		nama = nama + namaTambahan
		timData[idx].Nama = strings.TrimSpace(nama)
		fmt.Println("\nNama tim berhasil diupdate!")
	case 2:
		var fakultas string
		fmt.Print("Fakultas Baru: ")
		fmt.Scan(&fakultas)
		var fakultasTambahan string
		fmt.Scanln(&fakultasTambahan)
		fakultas = fakultas + fakultasTambahan
		timData[idx].Fakultas = strings.TrimSpace(fakultas)
		fmt.Println("\nFakultas berhasil diupdate!")
	case 3:
		var poin int
		fmt.Print("Poin Baru: ")
		fmt.Scan(&poin)
		timData[idx].Poin = poin
		fmt.Println("\nPoin berhasil diupdate!")
	default:
		fmt.Println("\nPilihan tidak valid.")
	}
}

// Hapus tim
func hapusTim() {
	var id string
	
	fmt.Println("\n--- HAPUS TIM ---")
	fmt.Print("Masukkan ID Tim yang akan dihapus: ")
	fmt.Scan(&id)
	
	idx := cariTimByID(id)
	if idx == -1 {
		fmt.Println("\nTim dengan ID tersebut tidak ditemukan.")
		return
	}
	
	tampilkanDetailTim(idx)
	
	var konfirmasi string
	fmt.Print("\nApakah Anda yakin ingin menghapus tim ini? (y/n): ")
	fmt.Scan(&konfirmasi)
	
	if strings.ToLower(konfirmasi) == "y" {
		timData[idx].IsDeleted = true
		fmt.Println("\nTim berhasil dihapus!")
	} else {
		fmt.Println("\nPenghapusan dibatalkan.")
	}
}

// Menampilkan statistik tim
func tampilkanStatistik() {
	if jumlahTim == 0 {
		fmt.Println("\nBelum ada tim yang terdaftar.")
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
		fmt.Println("\nBelum ada tim aktif yang terdaftar.")
		return
	}
	
	rataRataPoin := float64(totalPoin) / float64(jumlahTimAktif)
	
	fmt.Println("\n--- STATISTIK TIM ---")
	fmt.Println("Jumlah Tim Terdaftar :", jumlahTimAktif)
	fmt.Printf("Rata-rata Poin       : %.2f\n", rataRataPoin)
	fmt.Println("Poin Tertinggi       :", maxPoin, "(", timMaxPoin, ")")
	fmt.Println("Poin Terendah        :", minPoin, "(", timMinPoin, ")")
}

// Urutkan tim
func urutkanTim() {
	var pilihan, jenisUrutan int
	
	fmt.Println("\n--- URUTKAN TIM ---")
	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. ID Tim")
	fmt.Println("2. Nama Tim")
	fmt.Println("3. Fakultas")
	fmt.Println("4. Poin")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)
	
	fmt.Println("\nJenis pengurutan:")
	fmt.Println("1. Ascending (A-Z / 0-9)")
	fmt.Println("2. Descending (Z-A / 9-0)")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&jenisUrutan)
	
	var isAscending bool = (jenisUrutan == 1)
	var algoritma int
	
	fmt.Println("\nAlgoritma pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&algoritma)
	
	var timAktif []TimEsport
	for i := 0; i < jumlahTim; i++ {
		if !timData[i].IsDeleted {
			timAktif = append(timAktif, timData[i])
		}
	}
	
	if len(timAktif) == 0 {
		fmt.Println("\nBelum ada tim aktif yang terdaftar.")
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
			fmt.Println("\nPilihan tidak valid.")
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
			fmt.Println("\nPilihan tidak valid.")
			return
		}
	default:
		fmt.Println("\nPilihan tidak valid.")
		return
	}
	
	// Tampilkan hasil pengurutan
	fmt.Println("\n--- HASIL PENGURUTAN ---")
	fmt.Printf("%-5s %-20s %-20s %-5s\n", "ID", "Nama Tim", "Fakultas", "Poin")
	fmt.Println("---------------------------------------------------")
	
	for i := 0; i < len(timAktif); i++ {
		fmt.Printf("%-5s %-20s %-20s %-5d\n", 
			timAktif[i].ID, 
			timAktif[i].Nama, 
			timAktif[i].Fakultas, 
			timAktif[i].Poin)
	}
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