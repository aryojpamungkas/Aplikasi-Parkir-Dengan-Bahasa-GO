package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NMAX int = 1000

type User struct {
	Username string
	Password string
}

type tiketParkir struct {
	jenisKendaraan string
	PlatNomer      string
	waktuMasuk     string
	waktuKeluar    string
	jumlah         float64
}

type Admin struct {
	Username string
	Password string
}

var Pengguna User
var admin Admin
var tiket [NMAX]tiketParkir
var petugasParkir [NMAX]User
var tiketCount int
var petugasCount int

func main() {
	admin = Admin{"admin", "admin123"}

	fmt.Println("||================================================||")
	fmt.Println("||                                                ||")
	fmt.Println("||              Aplikasi Parkir                   ||")
	fmt.Println("||                   by :                         ||")
	fmt.Println("||           Galih Anggito Abimanyu               ||")
	fmt.Println("||            Aryo Jaty Pamungkas                 ||")
	fmt.Println("||     Tugas Besar Algoritma Pemrograman 2024     ||")
	fmt.Println("||                                                ||")
	fmt.Println("||================================================||")
	fmt.Println("*** Menu Utama ***")
	fmt.Println("1. Login")
	fmt.Println("2. Keluar")
	fmt.Println("------------------")

	var pilihan int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		login()
	case 2:
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid.")
		main()
	}
}

func login() {
	var username, password string
	fmt.Println("*** Login ***")
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if username == admin.Username && password == admin.Password {
		fmt.Println("Login berhasil sebagai admin.")
		adminMenu()
	} else {
		index := cariPetugasParkir(username)
		if index != -1 && petugasParkir[index].Password == password {
			Pengguna = petugasParkir[index]
			fmt.Println("Login berhasil sebagai petugas tiket parkir.")
			petugasMenu()
		} else {
			fmt.Println("Username atau Password salah.")
			var input string
			fmt.Println("Kembali ke menu awal ?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Print("pilihan: ")
			fmt.Scan(&input)
			if input == "1" {
				main()
			}
		}
	}
}

func adminMenu() {
	fmt.Println("*** Menu Admin ***")
	fmt.Println("1. Tambah Petugas Parkir")
	fmt.Println("2. Lihat Daftar Petugas Parkir")
	fmt.Println("3. Hapus Petugas Parkir")
	fmt.Println("4. Lihat Daftar Kendaraan dan Pendapatan")
	fmt.Println("5. Kembali ke Menu Utama")
	fmt.Println("------------------")

	var pilihan int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tambahPetugasParkir()
	case 2:
		lihatDaftarPetugasParkir()
	case 3:
		HapusPetugasParkir()
	case 4:
		LihatDaftarKendaraanDanPendapatan()
	case 5:
		main()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func tambahPetugasParkir() {
	fmt.Println("*** Tambah Petugas Parkir ***")
	var username, password string

	fmt.Print("Masukkan username baru: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password baru: ")
	fmt.Scanln(&password)

	if petugasCount < NMAX {
		petugasParkir[petugasCount] = User{Username: username, Password: password}
		petugasCount++
		sortPetugasParkir() // Urutkan setelah menambahkan petugas baru
		fmt.Println("Petugas parkir dengan username", username, "telah ditambahkan.")
	} else {
		fmt.Println("Maksimal petugas parkir sudah tercapai.")
	}

	var input string
	fmt.Println("Press x to continue...")
	fmt.Scan(&input)
	if input == "x" {
		main()
	}
}

func lihatDaftarPetugasParkir() {
	fmt.Println("*** Daftar Petugas Parkir ***")
	if petugasCount == 0 {
		fmt.Println("Belum ada petugas parkir terdaftar.")
	} else {
		for i := 0; i < petugasCount; i++ {
			fmt.Printf("%d. Username: %s\n", i+1, petugasParkir[i].Username)
		}
	}
	fmt.Println("------------------")
	var input string
	fmt.Println("Press x to continue...")
	fmt.Scan(&input)
	if input == "x" {
		main()
	}

}

func HapusPetugasParkir() {
	fmt.Println("*** Hapus Petugas Parkir ***")
	fmt.Print("Masukkan username petugas parkir yang ingin dihapus: ")
	var username string
	fmt.Scanln(&username)

	index := cariPetugasParkir(username)
	if index != -1 {
		for i := index; i < petugasCount-1; i++ {
			petugasParkir[i] = petugasParkir[i+1]
		}
		petugasCount--
		fmt.Println("Petugas parkir dengan username", username, "telah dihapus.")
	} else {
		fmt.Println("Petugas parkir dengan username", username, "tidak ditemukan.")
	}

	var input string
	fmt.Println("Press x to continue...")
	fmt.Scan(&input)
	if input == "x" {
		main()
	}
}

func LihatDaftarKendaraanDanPendapatan() {
	fmt.Println("*** Daftar Kendaraan dan Pendapatan ***")
	if tiketCount == 0 {
		fmt.Println("Belum ada transaksi parkir.")
	} else {
		totalPendapatan := 0.0
		for i := 0; i < tiketCount; i++ {
			fmt.Println("Jenis Kendaraan:", tiket[i].jenisKendaraan)
			fmt.Println("Plat Nomer:", tiket[i].PlatNomer)
			fmt.Println("Waktu Masuk:", tiket[i].waktuMasuk)
			fmt.Println("Waktu Keluar:", tiket[i].waktuKeluar)
			fmt.Printf("Tarif Parkir: Rp.%.2f\n", tiket[i].jumlah)
			fmt.Println()
			totalPendapatan += tiket[i].jumlah
		}
		fmt.Printf("Total Pendapatan Hari Ini: Rp.%.2f\n", totalPendapatan)
	}
	fmt.Println("------------------")
	var input string
	fmt.Println("Press x to continue...")
	fmt.Scan(&input)
	if input == "x" {
		main()
	}
}

func petugasMenu() {
	fmt.Println("*** Menu Petugas ***")
	fmt.Println("1. Input Transaksi Masuk")
	fmt.Println("2. Input Transaksi Keluar")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Println("------------------")

	var pilihan int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		InputTransaksiMasuk()
	case 2:
		InputTransaksiKeluar()
	case 3:
		main()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func InputTransaksiMasuk() {
	var jenis, plat, waktuMasuk string

	fmt.Print("Masukkan jenis kendaraan (motor / mobil) : ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan nomor kendaraan : ")
	fmt.Scan(&plat)
	fmt.Print("Masukkan waktu masuk (Jam:Menit) : ")
	fmt.Scan(&waktuMasuk)

	for i := 0; i < tiketCount; i++ {
		if tiket[i].jenisKendaraan == jenis && tiket[i].PlatNomer == plat {
			fmt.Println("Kendaraan sudah terdaftar.")
		}
	}

	if tiketCount < NMAX {
		tiket[tiketCount] = tiketParkir{
			jenisKendaraan: jenis,
			PlatNomer:      plat,
			waktuMasuk:     waktuMasuk,
		}
		tiketCount++
		fmt.Println("Transaksi parkir masuk berhasil dicatat.")
	} else {
		fmt.Println("Kapasitas tiket parkir sudah penuh.")
	}

	var input string
	fmt.Println("Press x to continue...")
	fmt.Scan(&input)
	if input == "x" {
		main()
	}
}

func InputTransaksiKeluar() {
	var jenis, plat, waktuKeluar string
	var bayar, kembalian float64

	fmt.Print("Masukkan jenis kendaraan (motor / mobil) : ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan nomor kendaraan : ")
	fmt.Scan(&plat)
	fmt.Print("Masukkan waktu Keluar (Jam:Menit) : ")
	fmt.Scan(&waktuKeluar)

	temp := cariTiketParkir(jenis, plat)
	if temp != -1 {
		waktuMasuk := tiket[temp].waktuMasuk
		durasi := hitungDurasiParkir(waktuMasuk, waktuKeluar)
		var tarif float64
		if jenis == "mobil" {
			fmt.Println("60 menit pertama Rp.4000 ")
			fmt.Println("Setelahnya Rp.2000 per 60 menit berikutnya.")
			if durasi <= 60 {
				tarif = 4000
			} else {
				tarif = 4000 + float64(2000*((durasi-60)/60))
			}
		} else if jenis == "motor" {
			fmt.Println("60 menit pertama Rp.2000 ")
			fmt.Println("Setelahnya Rp.1000 per 60 menit berikutnya.")
			if durasi <= 60 {
				tarif = 2000
			} else {
				tarif = 2000 + float64(1000*((durasi-60)/60))
			}
		}

		tiket[temp].waktuKeluar = waktuKeluar
		tiket[temp].jumlah = tarif
		fmt.Println("Total tarif parkir:", tarif)
		fmt.Println()
		fmt.Print("Masukan uang yang dibayar: Rp. ")
		fmt.Scan(&bayar)
		kembalian = bayar - tarif
		fmt.Println("Total kembalian anda: Rp.", kembalian)
	} else {
		fmt.Println("Kendaraan tidak ditemukan.")
	}

	var input string
	fmt.Println("Press x to continue...")
	fmt.Scan(&input)
	if input == "x" {
		main()
	}
}

// Fungsi untuk menghitung durasi parkir dalam menit
func hitungDurasiParkir(waktuMasuk, waktuKeluar string) int {
	menitMasuk := ubahKeMenit(waktuMasuk)
	menitKeluar := ubahKeMenit(waktuKeluar)

	durasiParkir := menitKeluar - menitMasuk

	return durasiParkir
}

// Fungsi untuk mengubah format jam:menit menjadi menit
func ubahKeMenit(waktu string) int {
	hasil := strings.Split(waktu, ":")
	totalMenit := 0

	for i := 0; i < len(hasil); i++ {
		num, error := strconv.Atoi(hasil[i])
		if error != nil {
			fmt.Println(error)
			return 0
		}
		if i == 0 { // jam
			totalMenit += num * 60
		} else if i == 1 { //  menit
			totalMenit += num
		}
	}

	return totalMenit
}

// Fungsi untuk mencari petugas parkir dengan username
func cariPetugasParkir(username string) int {
	for i := 0; i < petugasCount; i++ {
		if petugasParkir[i].Username == username {
			return i
		}
	}
	return -1
}

// Fungsi untuk mengurutkan petugas parkir berdasarkan username
func sortPetugasParkir() {
	for i := 0; i < petugasCount-1; i++ {
		for j := 0; j < petugasCount-i-1; j++ {
			if petugasParkir[j].Username > petugasParkir[j+1].Username {
				petugasParkir[j], petugasParkir[j+1] = petugasParkir[j+1], petugasParkir[j]
			}
		}
	}
}

// Fungsi untuk mencari tiket parkir dengan jenis kendaraan dan plat nomer
func cariTiketParkir(jenis, plat string) int {
	for i := 0; i < tiketCount; i++ {
		if tiket[i].jenisKendaraan == jenis && tiket[i].PlatNomer == plat && tiket[i].waktuKeluar == "" {
			return i
		}
	}
	return -1
}
