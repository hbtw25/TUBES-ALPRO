package main

import (
	"fmt"
	"strings"
)

// Konstanta untuk kapasitas maksimum tim
const NMAX int = 5

// Struktur data untuk Tim Esport
type TimEsport struct {
	ID       string
	Nama     string
	Fakultas string
	Poin     int
}

// Variabel global untuk menyimpan data tim
var timData [NMAX]TimEsport // IS: Array kosong saat program dimulai (sebelum tambahDataDummy). FS: Berisi data tim selama runtime.
var jumlahTim int = 0       // IS: 0 saat program dimulai. FS: Jumlah aktual tim yang tersimpan.

// tambahDataDummy mengisi array timData dengan data awal.
// IS: timData mungkin kosong atau sudah berisi data, jumlahTim adalah jumlah tim saat ini.
// FS: Jika NMAX belum tercapai, timData ditambahkan hingga 3 tim dummy, jumlahTim diperbarui.
func tambahDataDummy() {
	if jumlahTim < NMAX {
		timData[jumlahTim] = TimEsport{ID: "T01", Nama: "Alpha Coders", Fakultas: "Informatika", Poin: 100}
		jumlahTim++
	}
	if jumlahTim < NMAX {
		timData[jumlahTim] = TimEsport{ID: "T02", Nama: "Beta Testers", Fakultas: "Elektro", Poin: 85}
		jumlahTim++
	}
	if jumlahTim < NMAX {
		timData[jumlahTim] = TimEsport{ID: "T03", Nama: "Gamma Gamers", Fakultas: "DKV", Poin: 92}
		jumlahTim++
	}
}

// main adalah fungsi utama yang menjalankan program.
// IS: Program baru dimulai, data dummy akan ditambahkan.
// FS: Program berakhir setelah pengguna memilih opsi keluar (0).
func main() {
	tambahDataDummy()

	var pilihan int
	padding := "                                  "

	// Loop utama program.
	// IS (setiap iterasi): Menunggu input pilihan dari pengguna.
	// FS (setiap iterasi): Pilihan pengguna diproses.
	// FS (loop): Berakhir ketika pilihan == 0.
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
		fmt.Scanln()
		var dummy string
		fmt.Scanln(&dummy)
	}
}

// tampilkanMenu menampilkan daftar opsi kepada pengguna.
// IS: Tidak ada state input khusus.
// FS: Menu tercetak di konsol.
func tampilkanMenu() {
	indent := "                                  "
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

// tambahTim menambahkan data tim baru.
// IS: jumlahTim < NMAX. Pengguna siap memasukkan data tim.
// FS: Jika ID unik & kapasitas ada, tim baru ditambahkan ke timData, jumlahTim++. Jika tidak, pesan error.
func tambahTim() {
	if jumlahTim >= NMAX {
		indent := "                                   "
		fmt.Println(indent + "Maaf, kapasitas tim sudah penuh!")
		return
	}

	var tim TimEsport
	indent := "                                  "
	fmt.Println(indent + "TAMBAH TIM BARU")

	// Loop validasi ID Tim.
	// IS (loop): ID tim belum tervalidasi.
	// FS (loop): tim.ID valid (unik).
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

// tampilkanSemuaTim menampilkan semua data tim yang tersimpan.
// IS: timData dan jumlahTim mencerminkan state tim saat ini.
// FS: Data semua tim tercetak, atau pesan jika tidak ada tim.
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
	// Loop tampilkan tim.
	// IS (loop): i = 0.
	// FS (loop): Semua tim (0 hingga jumlahTim-1) ditampilkan.
	for i := 0; i < jumlahTim; i++ {
		fmt.Printf(padding+"║ %-3s | %-20s | %-20s | %-13d ║\n",
			timData[i].ID, timData[i].Nama, timData[i].Fakultas, timData[i].Poin)
	}
	fmt.Println(padding + "╚═══════════════════════════════════════════════════════════════════╝")
}

// sequentialSearchByIDGlobal mencari tim berdasarkan ID.
// IS: id = ID yang dicari. timData, jumlahTim adalah state global.
// FS: Index i jika timData[i].ID == id, atau -1 jika tidak ditemukan.
func sequentialSearchByIDGlobal(id string) int {
	var i int = 0
	var found bool = false
	// Loop pencarian.
	// IS (loop): i = 0, found = false.
	// FS (loop): (found == true && timData[i].ID == id) || (i == jumlahTim && found == false).
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

// sequentialSearchByNamaGlobal mencari tim berdasarkan Nama (case-insensitive).
// IS: nama = Nama yang dicari. timData, jumlahTim adalah state global.
// FS: Index i jika nama tim cocok (case-insensitive), atau -1 jika tidak.
func sequentialSearchByNamaGlobal(nama string) int {
	var i int = 0
	var found bool = false
	namaLower := strings.ToLower(nama)
	// Loop pencarian.
	// IS (loop): i = 0, found = false.
	// FS (loop): (found == true && cocok nama) || (i == jumlahTim && found == false).
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

// binarySearchByID mencari tim berdasarkan ID dalam array `arr` (sebanyak `n` elemen, terurut).
// Precondition: arr[0...n-1] terurut berdasarkan ID.
// IS: arr menunjuk ke array, n jumlah elemen, idCari adalah target. kr=0, kn=n-1, found=false.
// FS: Index med jika arr[med].ID == idCari, atau -1 jika tidak.
func binarySearchByID(arr *[NMAX]TimEsport, n int, idCari string) int {
	var kr, kn, med int
	kr = 0
	kn = n - 1
	var found bool = false
	// Loop pencarian biner.
	// IS (loop): kr <= kn, found = false. Batas pencarian [kr...kn].
	// FS (loop): (found == true && arr[med].ID == idCari) || (kr > kn && found == false).
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
	if found {
		return med
	}
	return -1
}

// binarySearchByNama mencari tim berdasarkan Nama dalam array `arr` (sebanyak `n` elemen, terurut, case-insensitive).
// Precondition: arr[0...n-1] terurut berdasarkan Nama (case-insensitive).
// IS: arr menunjuk ke array, n jumlah elemen, namaCari target. kr=0, kn=n-1, found=false.
// FS: Index med jika nama cocok (case-insensitive), atau -1 jika tidak.
func binarySearchByNama(arr *[NMAX]TimEsport, n int, namaCari string) int {
	var kr, kn, med int
	kr = 0
	kn = n - 1
	var found bool = false
	namaCariLower := strings.ToLower(namaCari)
	// Loop pencarian biner.
	// IS (loop): kr <= kn, found = false. Batas [kr...kn].
	// FS (loop): (found == true && cocok nama) || (kr > kn && found == false).
	for kr <= kn && !found {
		med = (kr + kn) / 2
		namaDataMedLower := strings.ToLower(arr[med].Nama)
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

// cariTim memproses pencarian tim.
// IS: Pengguna memilih opsi "Cari Tim".
// FS: Hasil pencarian ditampilkan.
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

	if algorithmChoice < 1 || algorithmChoice > 2 {
		fmt.Println(padding + "Pilihan algoritma tidak valid.")
		return
	}

	switch PilihanPencarian {
	case 1: // Cari berdasarkan ID
		var id string
		fmt.Print(padding + "Masukkan ID Tim: ")
		fmt.Scan(&id)
		if algorithmChoice == 1 {
			idx := sequentialSearchByIDGlobal(id)
			if idx != -1 {
				tampilkanDetailTimByIndex(idx)
			} else {
				fmt.Println(padding + "Tim dengan ID tersebut tidak ditemukan (Sequential).")
			}
		} else if algorithmChoice == 2 {
			if jumlahTim == 0 {
				fmt.Println(padding + "Tidak ada data untuk dilakukan binary search.")
				return
			}
			var searchableTeamsArray [NMAX]TimEsport                                   // IS: Array kosong.
			copy(searchableTeamsArray[:jumlahTim], timData[:jumlahTim])                // FS: Array berisi salinan data tim.
			selectionSortByID(&searchableTeamsArray, jumlahTim, true)                  // IS: Array belum terurut. FS: Array terurut by ID.
			idxInSortedArray := binarySearchByID(&searchableTeamsArray, jumlahTim, id) // IS: Array terurut. FS: idxInSortedArray berisi hasil.
			if idxInSortedArray != -1 {
				tampilkanDetailTimData(searchableTeamsArray[idxInSortedArray])
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
		if algorithmChoice == 1 {
			idx := sequentialSearchByNamaGlobal(nama)
			if idx != -1 {
				tampilkanDetailTimByIndex(idx)
			} else {
				fmt.Println(padding + "Tim dengan nama tersebut tidak ditemukan (Sequential).")
			}
		} else if algorithmChoice == 2 {
			if jumlahTim == 0 {
				fmt.Println(padding + "Tidak ada data untuk dilakukan binary search.")
				return
			}
			var searchableTeamsArray [NMAX]TimEsport                                       // IS: Array kosong.
			copy(searchableTeamsArray[:jumlahTim], timData[:jumlahTim])                    // FS: Array berisi salinan data tim.
			selectionSortByNama(&searchableTeamsArray, jumlahTim, true)                    // IS: Array belum terurut. FS: Array terurut by Nama.
			idxInSortedArray := binarySearchByNama(&searchableTeamsArray, jumlahTim, nama) // IS: Array terurut. FS: idxInSortedArray berisi hasil.
			if idxInSortedArray != -1 {
				tampilkanDetailTimData(searchableTeamsArray[idxInSortedArray])
			} else {
				fmt.Println(padding + "Tim dengan nama tersebut tidak ditemukan (Binary).")
			}
		}
	}
}

// tampilkanDetailTimByIndex menampilkan detail tim dari timData berdasarkan indeks.
// IS: idx adalah indeks valid di timData (0 <= idx < jumlahTim).
// FS: Detail tim timData[idx] tercetak.
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

// tampilkanDetailTimData menampilkan detail tim dari struct TimEsport.
// IS: tim berisi data tim yang akan ditampilkan.
// FS: Detail tim tercetak.
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

// updateTim memperbarui data tim.
// IS: Pengguna memasukkan ID tim.
// FS: Jika tim ditemukan, field yang dipilih diperbarui di timData. Jika tidak, pesan error.
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
		timData[idx].Nama = strings.TrimSpace(nama) // FS: timData[idx].Nama diperbarui.
		fmt.Println(padding + "Nama tim berhasil diupdate!")
	case 2:
		var fakultas string
		fmt.Print(padding + "Fakultas Baru: ")
		fmt.Scan(&fakultas)
		var fakultasTambahan string
		fmt.Scanln(&fakultasTambahan)
		fakultas = fakultas + fakultasTambahan
		timData[idx].Fakultas = strings.TrimSpace(fakultas) // FS: timData[idx].Fakultas diperbarui.
		fmt.Println(padding + "Fakultas berhasil diupdate!")
	case 3:
		var poin int
		fmt.Print(padding + "Poin Baru: ")
		fmt.Scan(&poin)
		timData[idx].Poin = poin // FS: timData[idx].Poin diperbarui.
		fmt.Println(padding + "Poin berhasil diupdate!")
	default:
		fmt.Println(padding + "Pilihan tidak valid.")
	}
}

// hapusTim menghapus data tim.
// IS: Pengguna memasukkan ID tim.
// FS: Jika tim ditemukan & dikonfirmasi, tim dihapus dari timData, jumlahTim--. Jika tidak, pesan error/batal.
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
		// Loop geser elemen.
		// IS (loop): i = idx.
		// FS (loop): Elemen timData[idx+1...jumlahTim-1] digeser ke kiri.
		for i := idx; i < jumlahTim-1; i++ {
			timData[i] = timData[i+1]
		}
		jumlahTim--
		timData[jumlahTim] = TimEsport{} // Mengosongkan posisi terakhir yang ditinggalkan.
		fmt.Println(padding + "Tim berhasil dihapus!")
	} else {
		fmt.Println(padding + "Penghapusan dibatalkan.")
	}
}

// tampilkanStatistik menghitung dan menampilkan statistik tim.
// IS: timData dan jumlahTim mencerminkan state tim saat ini.
// FS: Statistik tim tercetak. Pesan jika tidak ada tim.
func tampilkanStatistik() {
	padding := "                                  "
	if jumlahTim == 0 {
		fmt.Println(padding + "Belum ada tim yang terdaftar.")
		return
	}

	var totalPoin int = 0
	var minPoin int
	var maxPoin int
	var timMinPoin string
	var timMaxPoin string

	// Loop hitung statistik.
	// IS (loop): i = 0. Variabel statistik akan diinisialisasi/diperbarui.
	// FS (loop): totalPoin, minPoin, maxPoin, timMinPoin, timMaxPoin berisi nilai yang sesuai.
	for i := 0; i < jumlahTim; i++ {
		totalPoin += timData[i].Poin
		if i == 0 {
			minPoin = timData[i].Poin
			maxPoin = timData[i].Poin
			timMinPoin = timData[i].Nama
			timMaxPoin = timData[i].Nama
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

// selectionSortByID mengurutkan `n` elemen pertama dari `arr` berdasarkan ID.
// IS: arr menunjuk ke array, n jumlah elemen valid. isAscending menentukan urutan.
// FS: n elemen pertama arr terurut berdasarkan ID sesuai isAscending.
func selectionSortByID(arr *[NMAX]TimEsport, n int, isAscending bool) {
	var i, j, idx_min_max int
	var t TimEsport
	// Loop utama Selection Sort (pass).
	// IS (outer loop): i = 1. arr[0...i-2] terurut (awalnya kosong).
	// FS (outer loop): arr[0...n-1] terurut.
	i = 1
	for i <= n-1 {
		// IS (pass i): arr[0...i-2] terurut. arr[i-1...n-1] belum tentu. idx_min_max = i-1 (kandidat awal).
		idx_min_max = i - 1
		// Inner loop: cari elemen min/max di sisa array.
		// IS (inner loop): j = i. idx_min_max adalah kandidat min/max dari arr[i-1].
		// FS (inner loop): idx_min_max menunjuk ke elemen min/max sebenarnya di arr[i-1...n-1].
		j = i
		for j < n {
			if isAscending {
				if arr[j].ID < arr[idx_min_max].ID {
					idx_min_max = j
				}
			} else {
				if arr[j].ID > arr[idx_min_max].ID {
					idx_min_max = j
				}
			}
			j = j + 1
		}
		t = arr[idx_min_max]
		arr[idx_min_max] = arr[i-1]
		arr[i-1] = t
		// FS (pass i): arr[0...i-1] terurut.
		i = i + 1
	}
}

// selectionSortByNama mengurutkan `n` elemen dari `arr` berdasarkan Nama (case-insensitive).
// IS: arr, n, isAscending.
// FS: n elemen arr terurut berdasarkan Nama.
func selectionSortByNama(arr *[NMAX]TimEsport, n int, isAscending bool) {
	var i, j, idx_min_max int
	var t TimEsport
	i = 1
	for i <= n-1 {
		idx_min_max = i - 1
		j = i
		for j < n {
			namaJLower := strings.ToLower(arr[j].Nama)
			namaIdxMinMaxLower := strings.ToLower(arr[idx_min_max].Nama)
			if isAscending {
				if namaJLower < namaIdxMinMaxLower {
					idx_min_max = j
				}
			} else {
				if namaJLower > namaIdxMinMaxLower {
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

// selectionSortByFakultas mengurutkan `n` elemen dari `arr` berdasarkan Fakultas (case-insensitive).
// IS: arr, n, isAscending.
// FS: n elemen arr terurut berdasarkan Fakultas.
func selectionSortByFakultas(arr *[NMAX]TimEsport, n int, isAscending bool) {
	var i, j, idx_min_max int
	var t TimEsport
	i = 1
	for i <= n-1 {
		idx_min_max = i - 1
		j = i
		for j < n {
			fakultasJLower := strings.ToLower(arr[j].Fakultas)
			fakultasIdxMinMaxLower := strings.ToLower(arr[idx_min_max].Fakultas)
			if isAscending {
				if fakultasJLower < fakultasIdxMinMaxLower {
					idx_min_max = j
				}
			} else {
				if fakultasJLower > fakultasIdxMinMaxLower {
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

// selectionSortByPoin mengurutkan `n` elemen dari `arr` berdasarkan Poin.
// IS: arr, n, isAscending.
// FS: n elemen arr terurut berdasarkan Poin.
func selectionSortByPoin(arr *[NMAX]TimEsport, n int, isAscending bool) {
	var i, j, idx_min_max int
	var t TimEsport
	i = 1
	for i <= n-1 {
		idx_min_max = i - 1
		j = i
		for j < n {
			if isAscending {
				if arr[j].Poin < arr[idx_min_max].Poin {
					idx_min_max = j
				}
			} else {
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

// insertionSortByID mengurutkan `n` elemen dari `arr` berdasarkan ID.
// IS: arr, n, isAscending.
// FS: n elemen arr terurut berdasarkan ID.
func insertionSortByID(arr *[NMAX]TimEsport, n int, isAscending bool) {
	var i, j int
	var temp TimEsport
	// Loop utama Insertion Sort.
	// IS (outer loop): i = 1. arr[0...i-1] terurut (awalnya arr[0]). temp = arr[i].
	// FS (outer loop): arr[0...n-1] terurut.
	for i = 1; i < n; i++ {
		temp = arr[i]
		j = i - 1
		// Inner loop: geser elemen di bagian terurut.
		// IS (inner loop): j >= 0. temp dibandingkan dengan arr[j].
		// FS (inner loop): Elemen digeser, j adalah posisi sebelum penyisipan temp atau j < 0.
		if isAscending {
			for j >= 0 && arr[j].ID > temp.ID {
				arr[j+1] = arr[j]
				j = j - 1
			}
		} else {
			for j >= 0 && arr[j].ID < temp.ID {
				arr[j+1] = arr[j]
				j = j - 1
			}
		}
		arr[j+1] = temp // FS (iterasi i): arr[0...i] terurut.
	}
}

// insertionSortByNama mengurutkan `n` elemen dari `arr` berdasarkan Nama (case-insensitive).
// IS: arr, n, isAscending.
// FS: n elemen arr terurut berdasarkan Nama.
func insertionSortByNama(arr *[NMAX]TimEsport, n int, isAscending bool) {
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
		} else {
			for j >= 0 && strings.ToLower(arr[j].Nama) < namatempLower {
				arr[j+1] = arr[j]
				j = j - 1
			}
		}
		arr[j+1] = temp
	}
}

// insertionSortByFakultas mengurutkan `n` elemen dari `arr` berdasarkan Fakultas (case-insensitive).
// IS: arr, n, isAscending.
// FS: n elemen arr terurut berdasarkan Fakultas.
func insertionSortByFakultas(arr *[NMAX]TimEsport, n int, isAscending bool) {
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
		} else {
			for j >= 0 && strings.ToLower(arr[j].Fakultas) < fakultastempLower {
				arr[j+1] = arr[j]
				j = j - 1
			}
		}
		arr[j+1] = temp
	}
}

// insertionSortByPoin mengurutkan `n` elemen dari `arr` berdasarkan Poin.
// IS: arr, n, isAscending.
// FS: n elemen arr terurut berdasarkan Poin.
func insertionSortByPoin(arr *[NMAX]TimEsport, n int, isAscending bool) {
	var i, j int
	var temp TimEsport
	for i = 1; i < n; i++ {
		temp = arr[i]
		j = i - 1
		if isAscending {
			for j >= 0 && arr[j].Poin > temp.Poin {
				arr[j+1] = arr[j]
				j = j - 1
			}
		} else {
			for j >= 0 && arr[j].Poin < temp.Poin {
				arr[j+1] = arr[j]
				j = j - 1
			}
		}
		arr[j+1] = temp
	}
}

// urutkanTim menampilkan data tim yang diurutkan.
// IS: Pengguna memilih "Urutkan Tim". timData berisi data tim.
// FS: Salinan data tim diurutkan dan ditampilkan. timData global tidak berubah.
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
	fmt.Println(padding + "║ 1. Ascending (A-Z / 0-9)                   ║")
	fmt.Println(padding + "║ 2. Descending (Z-A / 9-0)                  ║")
	fmt.Println(padding + "╚════════════════════════════════════════════╝")
	fmt.Print(padding + "Pilihan Anda: ")
	fmt.Scan(&jenisUrutan)
	if jenisUrutan < 1 || jenisUrutan > 2 {
		fmt.Println(padding + "Pilihan jenis pengurutan tidak valid.")
		return
	}
	isAscending := (jenisUrutan == 1)

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

	var timToSortArray [NMAX]TimEsport                    // IS: Array kosong.
	copy(timToSortArray[:jumlahTim], timData[:jumlahTim]) // FS: Array berisi salinan data tim.

	// IS (sebelum switch): timToSortArray adalah salinan data (belum terurut).
	// FS (setelah switch): timToSortArray terurut sesuai pilihan.
	switch algoritmaPilihan {
	case 1: // Selection Sort
		switch pilihanField {
		case 1:
			selectionSortByID(&timToSortArray, jumlahTim, isAscending)
		case 2:
			selectionSortByNama(&timToSortArray, jumlahTim, isAscending)
		case 3:
			selectionSortByFakultas(&timToSortArray, jumlahTim, isAscending)
		case 4:
			selectionSortByPoin(&timToSortArray, jumlahTim, isAscending)
		}
	case 2: // Insertion Sort
		switch pilihanField {
		case 1:
			insertionSortByID(&timToSortArray, jumlahTim, isAscending)
		case 2:
			insertionSortByNama(&timToSortArray, jumlahTim, isAscending)
		case 3:
			insertionSortByFakultas(&timToSortArray, jumlahTim, isAscending)
		case 4:
			insertionSortByPoin(&timToSortArray, jumlahTim, isAscending)
		}
	}

	fmt.Println()
	fmt.Println(padding + "╔═══════════════════════════════════════════════════════════════════╗")
	fmt.Println(padding + "║ ID  | Nama Tim             | Fakultas             | Poin          ║")
	fmt.Println(padding + "╠═══════════════════════════════════════════════════════════════════╣")
	// Loop tampilkan hasil urut.
	// IS (loop): i = 0. timToSortArray sudah terurut.
	// FS (loop): Semua elemen timToSortArray (hingga jumlahTim) ditampilkan.
	for i := 0; i < jumlahTim; i++ {
		fmt.Printf(padding+"║ %-3s | %-20s | %-20s | %-13d ║\n",
			timToSortArray[i].ID, timToSortArray[i].Nama, timToSortArray[i].Fakultas, timToSortArray[i].Poin)
	}
	fmt.Println(padding + "╚═══════════════════════════════════════════════════════════════════╝")
}
