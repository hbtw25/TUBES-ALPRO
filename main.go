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
var timData [5]TimEsport // Array statis dengan kapasitas 5 tim (sesuai file awal)
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
			fmt.Println(padding + "Pilihan tidak valid. Silakan coba lagi")
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
	fmt.Println(indent + "║                Tugas Besar Algortima Pemrograman Semester 2               ║")
	fmt.Println(indent + "║                                                                           ║")
	fmt.Println(indent + "║  Anggota:                                                                 ║")
	fmt.Println(indent + "║   - Harsya Brahmantyo Wibowo                                              ║")
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
	if jumlahTim >= len(timData) {
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

		if sequentialSearchByIDGlobal(tim.ID) == -1 {
			break
		}
		fmt.Println(indent + "ID Tim sudah digunakan. Silakan gunakan ID lain.")
	}

	fmt.Print(indent + "Nama Tim: ")
	fmt.Scan(&tim.Nama)
	var namaTambahan string
	fmt.Scanln(&namaTambahan)
	tim.Nama = tim.Nama + namaTambahan
	tim.Nama = strings.TrimSpace(tim.Nama)

	fmt.Print(indent + "Fakultas: ")
	fmt.Scan(&tim.Fakultas)
	var fakultasTambahan string
	fmt.Scanln(&fakultasTambahan)
	tim.Fakultas = tim.Fakultas + fakultasTambahan
	tim.Fakultas = strings.TrimSpace(tim.Fakultas)

	fmt.Print(indent + "Poin: ")
	fmt.Scan(&tim.Poin)

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

// --- Fungsi-Fungsi Pencarian ---

// sequentialSearchByIDGlobal mencari tim berdasarkan ID di array global timData.
// Mengacu pada Modul 11.1.
func sequentialSearchByIDGlobal(id string) int {
	var i int = 0
	var found bool = false // Status pencarian [cite: 1624]
	// Proses pencarian berhenti jika ditemukan atau sudah semua data dicek
	for i < jumlahTim && !found {
		if timData[i].ID == id {
			found = true
		} else {
			i++
		}
	}
	if found {
		return i // Mengembalikan indeks data yang ditemukan
	}
	return -1 // Data tidak ditemukan [cite: 1631]
}

// sequentialSearchByNamaGlobal mencari tim berdasarkan Nama di array global timData.
// Mengacu pada Modul 11.1 dan 11.2. [cite: 1632, 1633]
func sequentialSearchByNamaGlobal(nama string) int {
	var i int = 0
	var found bool = false
	namaLower := strings.ToLower(nama)
	for i < jumlahTim && !found {
		if strings.ToLower(timData[i].Nama) == namaLower {
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

// binarySearchByID mencari tim berdasarkan ID dalam slice yang sudah terurut.
// Mengacu pada Modul 12.1. [cite: 1672]
func binarySearchByID(slice []TimEsport, idCari string) int {
	var kr, kn, med int
	kr = 0              // Indeks kiri [cite: 1678]
	kn = len(slice) - 1 // Indeks kanan [cite: 1678]
	var found bool = false

	for kr <= kn && !found {
		med = (kr + kn) / 2 // Ambil data tengah [cite: 1673, 1678]
		if slice[med].ID == idCari {
			found = true
		} else if slice[med].ID < idCari { // Jika data tengah lebih kecil, geser ke kanan [cite: 1674, 1678]
			kr = med + 1
		} else { // Jika data tengah lebih besar, geser ke kiri [cite: 1676, 1678]
			kn = med - 1
		}
	}
	if found {
		return med // Kembalikan indeks elemen yang ditemukan [cite: 1678]
	}
	return -1 // Kembalikan -1 jika tidak ditemukan [cite: 1684]
}

// binarySearchByNama mencari tim berdasarkan Nama dalam slice yang sudah terurut.
// Mengacu pada Modul 12.1 dan 12.2. [cite: 1672, 1686]
func binarySearchByNama(slice []TimEsport, namaCari string) int {
	var kr, kn, med int
	kr = 0
	kn = len(slice) - 1
	var found bool = false
	namaCariLower := strings.ToLower(namaCari)

	for kr <= kn && !found {
		med = (kr + kn) / 2
		namaDataMedLower := strings.ToLower(slice[med].Nama)
		if namaDataMedLower == namaCariLower {
			found = true
		} else if namaDataMedLower < namaCariLower {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}
	if found {
		return med
	}
	return -1
}

// Mencari tim (menu)
func cariTim() {
	var PilihanPencarian int
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
	fmt.Scan(&PilihanPencarian)

	// Validasi untuk pilihan field pencarian
	if PilihanPencarian < 1 || PilihanPencarian > 2 {
		fmt.Println(padding + "Pilihan field pencarian tidak valid.")
		return
	}

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

	// Validasi untuk pilihan algoritma
	if algorithmChoice < 1 || algorithmChoice > 2 {
		fmt.Println(padding + "Pilihan algoritma tidak valid.")
		return
	}

	switch PilihanPencarian {
	case 1: // Cari berdasarkan ID
		var id string
		fmt.Print(padding + "Masukkan ID Tim: ")
		fmt.Scan(&id)

		if algorithmChoice == 1 { // Sequential Search berdasarkan ID
			idx := sequentialSearchByIDGlobal(id)
			if idx != -1 {
				tampilkanDetailTimByIndex(idx)
			} else {
				fmt.Println(padding + "Tim dengan ID tersebut tidak ditemukan (Sequential).")
			}
		} else if algorithmChoice == 2 { // Binary Search berdasarkan ID
			if jumlahTim == 0 {
				fmt.Println(padding + "Tidak ada data untuk dilakukan binary search.")
				return
			}
			searchableTeams := make([]TimEsport, jumlahTim)
			copy(searchableTeams, timData[:jumlahTim])
			// Menggunakan gaya selection sort dari modul (indeks 0 to n-1)
			// Untuk binary search, data harus diurutkan menaik
			selectionSortByID(searchableTeams, true) // true berarti urutan menaik

			idxInSortedSlice := binarySearchByID(searchableTeams, id)
			if idxInSortedSlice != -1 {
				tampilkanDetailTimData(searchableTeams[idxInSortedSlice])
			} else {
				fmt.Println(padding + "Tim dengan ID tersebut tidak ditemukan (Binary).")
			}
		}

	case 2: // Cari berdasarkan Nama
		var nama string
		fmt.Print(padding + "Masukkan Nama Tim: ")
		fmt.Scan(&nama)
		var namaTambahan string
		fmt.Scanln(&namaTambahan)
		nama = nama + namaTambahan
		nama = strings.TrimSpace(nama)

		if algorithmChoice == 1 { // Sequential Search berdasarkan Nama
			idx := sequentialSearchByNamaGlobal(nama)
			if idx != -1 {
				tampilkanDetailTimByIndex(idx)
			} else {
				fmt.Println(padding + "Tim dengan nama tersebut tidak ditemukan (Sequential).")
			}
		} else if algorithmChoice == 2 { // Binary Search berdasarkan Nama
			if jumlahTim == 0 {
				fmt.Println(padding + "Tidak ada data untuk dilakukan binary search.")
				return
			}
			searchableTeams := make([]TimEsport, jumlahTim)
			copy(searchableTeams, timData[:jumlahTim])
			selectionSortByNama(searchableTeams, true) // Urutkan dulu untuk Binary Search (menaik)

			idxInSortedSlice := binarySearchByNama(searchableTeams, nama)
			if idxInSortedSlice != -1 {
				tampilkanDetailTimData(searchableTeams[idxInSortedSlice])
			} else {
				fmt.Println(padding + "Tim dengan nama tersebut tidak ditemukan (Binary).")
			}
		}
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

// Menampilkan detail tim dari struct TimEsport (untuk hasil binary search dari slice copy)
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

	idx := sequentialSearchByIDGlobal(id)
	if idx == -1 {
		fmt.Println(padding + "Tim dengan ID tersebut tidak ditemukan.")
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

// Hapus tim (Penghapusan Fisik)
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

	idx := sequentialSearchByIDGlobal(id)
	if idx == -1 {
		fmt.Println(padding + "Tim dengan ID tersebut tidak ditemukan.")
		return
	}

	tampilkanDetailTimByIndex(idx)

	var konfirmasi string
	fmt.Print(padding + "Apakah Anda yakin ingin menghapus tim ini? (y/n): ")
	fmt.Scan(&konfirmasi)

	if strings.ToLower(konfirmasi) == "y" {
		for i := idx; i < jumlahTim-1; i++ {
			timData[i] = timData[i+1]
		}
		jumlahTim--
		timData[jumlahTim] = TimEsport{} // Mengosongkan data di posisi terakhir
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

	for i := 1; i < jumlahTim; i++ {
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

// --- Fungsi-Fungsi Pengurutan (disesuaikan dengan gaya modul) ---

// selectionSortByID mengurutkan slice TimEsport berdasarkan ID.
// Mengacu pada Modul 13.1 & 13.2. [cite: 1727, 1733]
// Indeks array Go dimulai dari 0, jadi loop disesuaikan.
// Modul seringkali menggunakan i dari 1 to n-1 atau 0 to n-1.
// Di sini i adalah iterasi luar (pass), j adalah loop dalam untuk pencarian minimum/maksimum.
func selectionSortByID(arr []TimEsport, isAscending bool) {
	n := len(arr)
	// i: iterasi luar (pass), j: loop dalam, idx_min_max: indeks nilai ekstrim
	var i, j, idx_min_max int
	var t TimEsport // Untuk pertukaran

	i = 1          // Iterasi luar dimulai dari 1 (elemen pertama arr[0] adalah acuan awal)
	for i <= n-1 { // Loop sebanyak n-1 iterasi luar
		idx_min_max = i - 1 // Inisialisasi idx_min_max ke elemen pertama dari bagian belum terurut
		// (arr[i-1] dalam iterasi ke-i)
		j = i // Loop dalam mulai dari elemen setelah arr[i-1]
		for j < n {
			if isAscending {
				if arr[j].ID < arr[idx_min_max].ID {
					idx_min_max = j
				}
			} else { // Urutan menurun
				if arr[j].ID > arr[idx_min_max].ID {
					idx_min_max = j
				}
			}
			j = j + 1
		}
		// Tukar elemen arr[i-1] dengan elemen pada arr[idx_min_max]
		t = arr[idx_min_max]
		arr[idx_min_max] = arr[i-1]
		arr[i-1] = t
		i = i + 1
	}
}

// selectionSortByNama mengurutkan slice TimEsport berdasarkan Nama.
func selectionSortByNama(arr []TimEsport, isAscending bool) {
	n := len(arr)
	// i: iterasi luar (pass), j: loop dalam, idx_min_max: indeks nilai ekstrim
	var i, j, idx_min_max int
	var t TimEsport // Untuk pertukaran
	i = 1
	for i <= n-1 {
		idx_min_max = i - 1
		j = i
		for j < n {
			if isAscending {
				if strings.ToLower(arr[j].Nama) < strings.ToLower(arr[idx_min_max].Nama) {
					idx_min_max = j
				}
			} else { // Urutan menurun
				if strings.ToLower(arr[j].Nama) > strings.ToLower(arr[idx_min_max].Nama) {
					idx_min_max = j
				}
			}
			j = j + 1
		}
		t = arr[idx_min_max]
		arr[idx_min_max] = arr[i-1]
		arr[i-1] = t
		i = i + 1
	}
}

// selectionSortByFakultas mengurutkan slice TimEsport berdasarkan Fakultas.
func selectionSortByFakultas(arr []TimEsport, isAscending bool) {
	n := len(arr)
	// i: iterasi luar (pass), j: loop dalam, idx_min_max: indeks nilai ekstrim
	var i, j, idx_min_max int
	var t TimEsport // Untuk pertukaran
	i = 1
	for i <= n-1 {
		idx_min_max = i - 1
		j = i
		for j < n {
			if isAscending {
				if strings.ToLower(arr[j].Fakultas) < strings.ToLower(arr[idx_min_max].Fakultas) {
					idx_min_max = j
				}
			} else { // Urutan menurun
				if strings.ToLower(arr[j].Fakultas) > strings.ToLower(arr[idx_min_max].Fakultas) {
					idx_min_max = j
				}
			}
			j = j + 1
		}
		t = arr[idx_min_max]
		arr[idx_min_max] = arr[i-1]
		arr[i-1] = t
		i = i + 1
	}
}

// selectionSortByPoin mengurutkan slice TimEsport berdasarkan Poin.
func selectionSortByPoin(arr []TimEsport, isAscending bool) {
	n := len(arr)
	// i: iterasi luar (pass), j: loop dalam, idx_min_max: indeks nilai ekstrim
	var i, j, idx_min_max int
	var t TimEsport // Untuk pertukaran
	i = 1
	for i <= n-1 {
		idx_min_max = i - 1
		j = i
		for j < n {
			if isAscending {
				if arr[j].Poin < arr[idx_min_max].Poin {
					idx_min_max = j
				}
			} else { // Urutan menurun
				if arr[j].Poin > arr[idx_min_max].Poin {
					idx_min_max = j
				}
			}
			j = j + 1
		}
		t = arr[idx_min_max]
		arr[idx_min_max] = arr[i-1]
		arr[i-1] = t
		i = i + 1
	}
}

// insertionSortByID mengurutkan slice TimEsport berdasarkan ID.
// Mengacu pada Modul 14.1 & 14.2. [cite: 1764, 1771]
// Modul menunjukkan loop i dari 1, j dari i, dan temp = a[j], lalu while j > 0.
// Di sini i adalah batas antara bagian terurut dan tidak terurut.
func insertionSortByID(arr []TimEsport, isAscending bool) {
	n := len(arr)
	var i, j int
	var temp TimEsport

	for i = 1; i < n; i++ {
		temp = arr[i] // Elemen yang akan disisipkan [cite: 1770]
		j = i - 1     // Indeks terakhir dari bagian yang sudah terurut

		if isAscending {
			// Geser elemen dari bagian terurut yang lebih besar dari temp
			for j >= 0 && arr[j].ID > temp.ID {
				arr[j+1] = arr[j] // Geser ke kanan [cite: 1770]
				j = j - 1
			}
		} else { // Urutan menurun
			for j >= 0 && arr[j].ID < temp.ID {
				arr[j+1] = arr[j]
				j = j - 1
			}
		}
		arr[j+1] = temp // Sisipkan temp pada posisi yang tepat [cite: 1770]
	}
}

// insertionSortByNama mengurutkan slice TimEsport berdasarkan Nama.
func insertionSortByNama(arr []TimEsport, isAscending bool) {
	n := len(arr)
	var i, j int
	var temp TimEsport

	for i = 1; i < n; i++ {
		temp = arr[i]
		j = i - 1
		namatempLower := strings.ToLower(temp.Nama)
		if isAscending {
			for j >= 0 && strings.ToLower(arr[j].Nama) > namatempLower {
				arr[j+1] = arr[j]
				j = j - 1
			}
		} else { // Urutan menurun
			for j >= 0 && strings.ToLower(arr[j].Nama) < namatempLower {
				arr[j+1] = arr[j]
				j = j - 1
			}
		}
		arr[j+1] = temp
	}
}

// insertionSortByFakultas mengurutkan slice TimEsport berdasarkan Fakultas.
func insertionSortByFakultas(arr []TimEsport, isAscending bool) {
	n := len(arr)
	var i, j int
	var temp TimEsport

	for i = 1; i < n; i++ {
		temp = arr[i]
		j = i - 1
		fakultastempLower := strings.ToLower(temp.Fakultas)
		if isAscending {
			for j >= 0 && strings.ToLower(arr[j].Fakultas) > fakultastempLower {
				arr[j+1] = arr[j]
				j = j - 1
			}
		} else { // Urutan menurun
			for j >= 0 && strings.ToLower(arr[j].Fakultas) < fakultastempLower {
				arr[j+1] = arr[j]
				j = j - 1
			}
		}
		arr[j+1] = temp
	}
}

// insertionSortByPoin mengurutkan slice TimEsport berdasarkan Poin.
// Mengacu pada Modul 14.1 & 14.2. [cite: 1764, 1771]
// Contoh modul menggunakan loop for i<=n-1; temp=T[j]; for j>0 && temp > T[j-1] untuk descending
// Di sini kita adaptasi untuk ascending/descending dengan parameter.
func insertionSortByPoin(arr []TimEsport, isAscending bool) {
	n := len(arr)
	var i, j int
	var temp TimEsport // Di modul disebut 'temp' [cite: 1770]

	for i = 1; i < n; i++ { // Loop i dari 1 hingga n-1 [cite: 1770]
		temp = arr[i] // Ambil elemen arr[i] sebagai temp (di modul arr[j] dengan j=i) [cite: 1770]
		j = i - 1     // j adalah indeks terakhir dari bagian yang sudah terurut

		if isAscending {
			// Geser elemen dari bagian terurut yang lebih besar dari temp.Poin
			for j >= 0 && arr[j].Poin > temp.Poin { // Kondisi while di modul [cite: 1770]
				arr[j+1] = arr[j] // Geser elemen ke kanan [cite: 1770]
				j = j - 1         // Pindah ke elemen sebelumnya di bagian terurut [cite: 1770]
			}
		} else { // Urutan menurun (sesuai contoh descending di modul [cite: 1770])
			for j >= 0 && arr[j].Poin < temp.Poin { // `temp > T[j-1]` di modul untuk descending
				arr[j+1] = arr[j]
				j = j - 1
			}
		}
		arr[j+1] = temp // Sisipkan temp pada posisi yang tepat [cite: 1770]
	}
}

// Urutkan tim
func urutkanTim() {
	var pilihanField, jenisUrutan, algoritmaPilihan int
	padding := "                                  "

	if jumlahTim == 0 {
		fmt.Println(padding + "Belum ada tim yang terdaftar untuk diurutkan.")
		return
	}

	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════╗")
	fmt.Println(padding + "║     URUTKAN TIM BERDASARKAN   ║")
	fmt.Println(padding + "╠═══════════════════════════════╣")
	fmt.Println(padding + "║ 1. ID Tim                     ║")
	fmt.Println(padding + "║ 2. Nama Tim                   ║")
	fmt.Println(padding + "║ 3. Fakultas                   ║")
	fmt.Println(padding + "║ 4. Poin                       ║")
	fmt.Println(padding + "╚═══════════════════════════════╝")
	fmt.Print(padding + "Pilihan Anda: ")
	fmt.Scan(&pilihanField)

	if pilihanField < 1 || pilihanField > 4 {
		fmt.Println(padding + "Pilihan field tidak valid.")
		return
	}

	fmt.Println()
	fmt.Println(padding + "╔════════════════════════════════════════════╗")
	fmt.Println(padding + "║             JENIS PENGURUTAN               ║")
	fmt.Println(padding + "╠════════════════════════════════════════════╣")
	fmt.Println(padding + "║ 1. Ascending (A-Z / 0-9)                   ║") // "Ascending" di sini adalah bagian dari teks UI
	fmt.Println(padding + "║ 2. Descending (Z-A / 9-0)                  ║") // "Descending" di sini adalah bagian dari teks UI
	fmt.Println(padding + "╚════════════════════════════════════════════╝")
	fmt.Print(padding + "Pilihan Anda: ")
	fmt.Scan(&jenisUrutan)

	if jenisUrutan < 1 || jenisUrutan > 2 {
		fmt.Println(padding + "Pilihan jenis pengurutan tidak valid.")
		return
	}
	isAscending := (jenisUrutan == 1) // true jika pilihan 1 (Ascending/Menaik)

	fmt.Println()
	fmt.Println(padding + "╔════════════════════════════════╗")
	fmt.Println(padding + "║       ALGORITMA PENGURUTAN     ║")
	fmt.Println(padding + "╠════════════════════════════════╣")
	fmt.Println(padding + "║ 1. Selection Sort              ║")
	fmt.Println(padding + "║ 2. Insertion Sort              ║")
	fmt.Println(padding + "╚════════════════════════════════╝")
	fmt.Print(padding + "Pilihan Anda: ")
	fmt.Scan(&algoritmaPilihan)

	if algoritmaPilihan < 1 || algoritmaPilihan > 2 {
		fmt.Println(padding + "Pilihan algoritma pengurutan tidak valid.")
		return
	}

	timToSort := make([]TimEsport, jumlahTim)
	copy(timToSort, timData[:jumlahTim])

	switch algoritmaPilihan {
	case 1: // Algoritma Selection Sort
		switch pilihanField {
		case 1:
			selectionSortByID(timToSort, isAscending)
		case 2:
			selectionSortByNama(timToSort, isAscending)
		case 3:
			selectionSortByFakultas(timToSort, isAscending)
		case 4:
			selectionSortByPoin(timToSort, isAscending)
		}
	case 2: // Algoritma Insertion Sort
		switch pilihanField {
		case 1:
			insertionSortByID(timToSort, isAscending)
		case 2:
			insertionSortByNama(timToSort, isAscending)
		case 3:
			insertionSortByFakultas(timToSort, isAscending)
		case 4:
			insertionSortByPoin(timToSort, isAscending)
		}
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
