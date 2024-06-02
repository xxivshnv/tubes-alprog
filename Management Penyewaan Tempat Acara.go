package main

import (
	"fmt"
)

const NMAX int = 50

type tabTempat [NMAX]iTempat
type iTempat struct {
	namaTempat        string
	lokasiTempat      string
	kapasitasMaksimum int
	fasilitas         []string
	hargaSewa         int
	riwayatSewa       []waktuSewa // int mewakilkan tanggal sewa
}

type waktuSewa struct {
	namaPenyewa string
	tanggal     int
	jamMulai    int
	jamSelesai  int
}

type tabUserAcc [NMAX]userAcc
type userAcc struct {
	userName  string
	password  string
	userClass string // "manajer" OR "pelanggan"
}

// FOR TESTING PURPOSE
func testingPurpose(dataTempat *tabTempat, dataUser *tabUserAcc, nDataTempat *int, nDataUser *int) {
	dataUser[0] = userAcc{
		userName:  "naruto",
		password:  "1234",
		userClass: "pelanggan",
	}
	dataUser[1] = userAcc{
		userName:  "sasuke",
		password:  "4444",
		userClass: "manajer",
	}
	*nDataUser = 2 // Corrected parameter
	dataTempat[0] = iTempat{
		namaTempat:        "hotel indonesia",
		lokasiTempat:      "jakarta",
		kapasitasMaksimum: 100,
		fasilitas:         []string{"Wifi", "AC"},
		hargaSewa:         1000000,
		riwayatSewa:       []waktuSewa{},
	}
	dataTempat[1] = iTempat{
		namaTempat:        "gelora bung karno",
		lokasiTempat:      "jakarta",
		kapasitasMaksimum: 200,
		fasilitas:         []string{"Proyektor", "AC"},
		hargaSewa:         2000000,
		riwayatSewa:       []waktuSewa{},
	}
	dataTempat[2] = iTempat{
		namaTempat:        "kebun raya bogor",
		lokasiTempat:      "bogor",
		kapasitasMaksimum: 400,
		fasilitas:         []string{"kantin", "sungai"},
		hargaSewa:         1500000,
		riwayatSewa:       []waktuSewa{},
	}
	*nDataTempat = 3
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////
////                                                                        LOGIN OR REGISTER FUNCTION
////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	var dataTempat tabTempat
	var nDataTempat int
	var dataUser tabUserAcc
	var nDataUser int

	//ADD DUMMY ACCOUNT FOR TESTING PURPOSE ONLY
	testingPurpose(&dataTempat, &dataUser, &nDataTempat, &nDataUser)

	var currentUserClass string // user class dari akun yang sedang dipakai "manajer" OR "pelanggan"
	var currentIndexAcc int     // index dari akun yang sedang dipakai

	mainMenu(&dataUser, &nDataUser, &currentUserClass, &currentIndexAcc)
}

// PUBLIC's FUNCTION START HERE

func mainMenu(dataUser *tabUserAcc, nDataUser *int, currentUserClass *string, currentIndexAcc *int) {
	/*
		 	IS : 	-
			FS : 	Mengembalikan string "login" atau "register" atau Keluar program sesuai dengan pilihan user ketika mengisi variable userChoice
	*/
	var userChoice int = 0
	for userChoice != 3 {
		fmt.Println("Masukan angka :")
		fmt.Println("1 untuk login")
		fmt.Println("2 untuk daftar")
		fmt.Println("3 untuk keluar program")
		fmt.Print("pilihan : ")
		fmt.Scan(&userChoice)
		fmt.Println("")
		if userChoice == 1 {
			login(*dataUser, *nDataUser, currentUserClass, currentIndexAcc)
			break
		} else if userChoice == 2 {
			register(dataUser, nDataUser, currentUserClass, currentIndexAcc)
			break
		} else if userChoice == 3 {
			*currentIndexAcc = -1
			*currentUserClass = "0"
			fmt.Print("Terimakasih Telah Menggunakan Program Sewa Tempat")
		} else {
			fmt.Println("\npilihan tidak tersedia, silahkan ulangi.")
		}
	}
	//Test the connected account
	if userChoice == 1 || userChoice == 2 {
		if *currentUserClass == "pelanggan" {
			menuPelanggan(*dataTempat, *nDataTempat, currentIndexAcc, currentUserClass)
		} else if *currentUserClass == "manajer" {
			menuManajer(dataTempat, nDataTempat, currentIndexAcc, currentUserClass)
		}
	}

}

func login(dataUser tabUserAcc, nDataUser int, currentUserClass *string, currentIndexAcc *int) {
	/*
		 	IS : 	variabel publik dataUser dan nDataUser yang menyimpan daftar akun yang sudah terdaftar di aplikasi dan jumlah akun yang sudah terdaftar di aplikasi
			     	alamat variabel public currentUserClass dan currentIndexAcc yang menyimpan user class dan index dari akun yang akan masuk
			FS : 	mengubah variabel publik currentUserClass dan currentIndexAcc sesuai dengan user class dan index dari akun yang berhasil masuk (login)
	*/
	var username, password string
	var found bool = false
	fmt.Println("Login. Silahkan masukan data.")
	fmt.Print("username : ")
	fmt.Scan(&username)
	fmt.Print("password : ")
	fmt.Scan(&password)
	// Loop untuk cek apakah ada akun dengan username dan password yang sesuai dengan input dari user
	for i := 0; i < nDataUser; i++ {
		if dataUser[i].userName == username && dataUser[i].password == password {
			*currentUserClass = dataUser[i].userClass
			*currentIndexAcc = i
			found = true
			fmt.Print("Akun berhasil login")
			break
		}
	}
	if found == false {
		fmt.Println("\n\nAkun tidak ditemukan. Silahkan ulangi.")
		login(dataUser, nDataUser, currentUserClass, currentIndexAcc)
	}
}

func register(dataUser *tabUserAcc, nDataUser *int, currentUserClass *string, currentIndexAcc *int) {
	/*
			IS :	alamat variabel publik dataUser dan nDataUser yang menyimpan daftar akun yang sudah terdaftar di aplikasi dan jumlah akun yang sudah terdaftar di aplikasi.
		 	     	alamat variabel public currentUserClass dan currentIndexAcc yang menyimpan user class dan index dari akun yang akan masuk.
		 	FS : 	memasukan akun baru dengan username dan password sesuai dengan masukan dari user kedalam variabel publik dataUser dan menambah nDataUser dengan 1
		  	     	mengubah variabel publik currentUserClass dan currentIndexAcc sesuai dengan user class dan index dari akun baru didaftarkan
	*/
	var username, password string
	var found bool = false
	fmt.Println("Register. Silahkan masukan data.")
	fmt.Print("username : ")
	fmt.Scan(&username)
	fmt.Print("password : ")
	fmt.Scan(&password)
	//Loop untuk cek apakah username sudah dipakai atau belum
	for i := 0; i < *nDataUser; i++ {
		if dataUser[i].userName == username {
			found = true
			break
		}
	}
	if found == true {
		fmt.Println("\n\nUsername sudah dipakai. Silahkan ulangi.")
		register(dataUser, nDataUser, currentUserClass, currentIndexAcc)
	} else {
		newUser := userAcc{
			userName:  username,
			password:  password,
			userClass: "pelanggan",
		}
		dataUser[*nDataUser] = newUser
		*nDataUser++
		*currentUserClass = newUser.userClass
		*currentIndexAcc = *nDataUser - 1
		fmt.Println("Akun berhasil register")
	}
}

func Logout(currentIndexAcc *int, currentUserClass *string) {
	/*
			IS :	alamat variabel public currentUserClass dan currentIndexAcc yang menyimpan user class dan index dari akun yang akan masuk.
			FS :	menyimpan -1 kedalam currentIndexAcc yang berarti tidak ada akun yang sedang masuk (login)
		 		menyimpan "0" kedalam currentUserClass yang berarti tidak ada user class
	*/
	*currentIndexAcc = -1
	*currentUserClass = "0"
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ///                                                                                                                                         ////
// //                                                       PUBLIC CLASS FUNCTION                                                              ////
// //                                                                                                                                          ////
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func cariTempat(dataTempat tabTempat, nDataTempat int) int {
	/*
			IS :	variabel dataTempat dan nDataTempat yang menyimpan datar tempat dan jumlah tempat
			FS : 	memanggil fungsi cariDenganNama atau cariDenganLokasi atau cariDenganTempat berdasarkan pilihan user dan menyimpan kedalam hasil fungsi kedalam variabel indexTempat
		 		mengembalikan nilai variabe indexTempat
	*/
	var userChoice int
	var indexTempat int = -1

	for indexTempat == -1 {
		fmt.Println("pilih metode pencarian. ketik : ")
		fmt.Println("1 untuk mencari dengan nama")
		fmt.Println("2 untuk mencari dengan lokasi")
		fmt.Println("3 untuk mencari dengan kapasitas")
		fmt.Print("masukan pilihan anda : ")
		fmt.Scan(&userChoice)
		fmt.Println("")

		switch userChoice { // ganti if else biasa
		case 1:
			indexTempat = cariDenganNama(dataTempat, nDataTempat)
		case 2:
			indexTempat = cariDenganLokasi(dataTempat, nDataTempat)
		case 3:
			indexTempat = cariDenganKapasitas(dataTempat, nDataTempat)
		default:
			fmt.Println("Pilihan tidak tersedia. silahkan pilih ulang.")
			return cariTempat(dataTempat, nDataTempat)
		}

		if indexTempat == -1 {
			fmt.Println("Tempat tidak ditemukan. silahkan pilih ulang.")
		}
	}
	return indexTempat

}

func cariDenganNama(dataTempat tabTempat, nDataTempat int) int {
	var namaYangDicari string
	fmt.Println("Pencarian.")
	fmt.Print("Masukan nama tempat: ")
	fmt.Scan(&namaYangDicari)
	fmt.Println("")
	for i := 0; i < nDataTempat; i++ {
		if namaYangDicari == dataTempat[i].namaTempat {
			return i
		}
	}
	return -1
}

func cariDenganLokasi(dataTempat tabTempat, nDataTempat int) int {
	var lokasiYangDicari string
	fmt.Println("Pencarian.")
	fmt.Print("Masukan lokasi tempat : ")
	fmt.Scan(&lokasiYangDicari)
	fmt.Println("")
	for i := 0; i < nDataTempat; i++ {
		if lokasiYangDicari == dataTempat[i].lokasiTempat {
			return i
		}
	}
	return -1
}

func cariDenganKapasitas(dataTempat tabTempat, nDataTempat int) int {
	var kapasitasYangDicari int
	fmt.Println("Pencarian.")
	fmt.Print("Masukan kapasitas tempat : ")
	fmt.Scan(&kapasitasYangDicari)
	fmt.Println("")
	for i := 0; i < nDataTempat; i++ {
		if kapasitasYangDicari == dataTempat[i].kapasitasMaksimum {
			return i
		}
	}
	return -1
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////
////                                                                    PELANGGAN CLASS FUNCTION
////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func menuPelanggan(dataTempat tabTempat, nDataTempat int, currentIndexAcc *int, currentUserClass *string) {
	/*

		IS : variabel data tempat,nDataTempat,currentindexAcc dan currentuserClass menyimpan data tempat, jumlah data tempat, index akun user saat ini dan class user
		FS : Memanggil fitur sewa tempat dan menginput data dari user kemudian memanggil isAvailabele untuk mengecek ketersediaan tempat,
		ketika kondisi true maka riwayat penyewaan user dari input sewa tempat akan ditampilkan

	*/
	var userChoice int = 0

	for userChoice != 2 {
		fmt.Println("Menu utama, Silahkan ketik")
		fmt.Println("1 untuk sewa tempat")
		fmt.Println("2 untuk keluar")
		fmt.Print("Pilihan : ")
		fmt.Scan(&userChoice)
		fmt.Println("")

		if userChoice == 1 {
			fiturSewaTempat(dataTempat, nDataTempat)
		} else if userChoice == 2 {
			Logout(currentIndexAcc, currentUserClass)
			fmt.Print("Berhasil keluar")
		} else {
			fmt.Println("Pilihan tidak tersedia, silahkan ulangi.")
			//menuPelanggan(dataTempat, nDataTempat, currentIndexAcc, currentUserClass) --> jika dipanggil dalam loop, tidak perlu dibuat rekrusif
		}
	}
}

func fiturSewaTempat(dataTempat tabTempat, nDataTempat int) {
	var currentIndexTempat int
	currentIndexTempat = cariTempat(dataTempat, nDataTempat)

	if currentIndexTempat == -1 {
		fmt.Println("Tempat tidak tersedia.")
		return
	}
	// Proses Booking Venue
	var booking waktuSewa
	fmt.Println("Masukkan nama penyewa: ")
	fmt.Scan(&booking.namaPenyewa) // nama penyewa diambil dari username akun yang sudah terlogin
	fmt.Println("Masukkan tanggal sewa (format: YYYYMMDD): ")
	fmt.Scan(&booking.tanggal)
	fmt.Println("Masukkan jam mulai (format 24 jam): ")
	fmt.Scan(&booking.jamMulai)
	fmt.Println("Masukkan jam selesai (format 24 jam): ")
	fmt.Scan(&booking.jamSelesai) // input nya durasi

	// Cek apakah tempat tersedia di waktu yang diminta
	if isAvailable(dataTempat[currentIndexTempat], booking) {
		dataTempat[currentIndexTempat].riwayatSewa = append(dataTempat[currentIndexTempat].riwayatSewa, booking)
		fmt.Println("Tempat berhasil disewa.")
	} else {
		fmt.Println("Tempat tidak tersedia pada waktu yang diminta.")
	}
}
func isAvailable(tempat iTempat, booking waktuSewa) bool {
	for _, sewa := range tempat.riwayatSewa {
		if sewa.tanggal == booking.tanggal {
			if (booking.jamMulai >= sewa.jamMulai && booking.jamMulai < sewa.jamSelesai) ||
				(booking.jamSelesai > sewa.jamMulai && booking.jamSelesai <= sewa.jamSelesai) ||
				(booking.jamMulai <= sewa.jamMulai && booking.jamSelesai >= sewa.jamSelesai) {
				return false
			}
		}
	}
	return true
}

func tampilkanRiwayatSewa(tempat iTempat) {
	fmt.Println("Riwayat Sewa untuk", tempat.namaTempat)
	for _, sewa := range tempat.riwayatSewa {
		fmt.Printf("Nama Penyewa: %s, Tanggal: %d, Jam Mulai: %d, Jam Selesai: %d\n",
			sewa.namaPenyewa, sewa.tanggal, sewa.jamMulai, sewa.jamSelesai)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////
////                                                      MANAJER CLASS FUNCTION
////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func menuManajer(dataTempat *tabTempat, nDataTempat *int, currentIndexAcc *int, currentUserClass *string) {
	/*

		IS : dataTempat berisi data tempat, nDataTempat berisi jumlah tempat, currentIndexAcc berisi index user saat ini
		currentUserClass berisi class dari user
		FS : opsi selanjutnya bergantung pada jawabab yang dipilih user antara menampilkan data terurut, merubah data tempat,
		menghapus data tempat atau logout dari menu



	*/
	var userChoice int = 0
	for userChoice != 4 {
		fmt.Println("Menu utama, Silahkan ketik")
		fmt.Println("1 untuk menampilkan data terurut")
		fmt.Println("2 untuk merubah data tempat")
		fmt.Println("3 untuk menghapus data tempat")
		fmt.Println("4 untuk keluar")
		fmt.Print("Pilihan : ")
		fmt.Scan(&userChoice)
		fmt.Println("")

		if userChoice == 1 {
			fiturUrutDataTempat(dataTempat, *nDataTempat)
			// urutDengaKapasitas(dataTempat, *nDataTempat)
			fiturTampilkanData(*dataTempat, *nDataTempat)
		} else if userChoice == 2 {
			fiturUbahDataTempat(dataTempat, *nDataTempat)
		} else if userChoice == 3 {
			fiturHapusDataTempat(dataTempat, nDataTempat)
		} else if userChoice == 4 {
			Logout(currentIndexAcc, currentUserClass)
			fmt.Print("Berhasil keluar")
		} else {
			fmt.Println("Pilian tidak tersedia, silahkan ulangi.")
		}
	}

}

func fiturUbahDataTempat(dataTempat *tabTempat, nDataTempat int) {
	/*

		IS : dataTempat berisi data tempat dan nDataTempat berisi jumlah data tempat
		FS :



	*/
	var indexTempat, userChoice int
	fmt.Println("Silahkan cari tempat terlebih dahulu")
	indexTempat = cariTempat(*dataTempat, nDataTempat)

	for userChoice < 1 || userChoice > 5 {
		userChoice = menuUbahDataTempat(*dataTempat, indexTempat)
	}
	fmt.Println("Data sebelum diubah")
	cetakDataTempat(*dataTempat, indexTempat)

	switch userChoice {
	case 1:
		UbahDataNamaTempat(dataTempat, indexTempat)
	case 2:
		UbahDataLokasiTempat(dataTempat, indexTempat)
	case 3:
		UbahDataKapasitasTempat(dataTempat, indexTempat)
	case 4:
		UbahDataFasilitasTempat(dataTempat, indexTempat)
	case 5:
		UbahDataHargaTempat(dataTempat, indexTempat)
	}
	fmt.Println("Data setelah diubah")
	cetakDataTempat(*dataTempat, indexTempat)
}

func menuUbahDataTempat(dataTempat tabTempat, indexTempat int) int {
	var userChoice int
	fmt.Printf("Apa yang ingin anda ubah dari data tempat %s? Ketik\n", dataTempat[indexTempat].namaTempat)
	fmt.Println("1 untuk ubah nama tempat")
	fmt.Println("2 untuk ubah lokasi tempat")
	fmt.Println("3 untuk ubah kapasitas maksimum")
	fmt.Println("4 untuk ubah fasilitas")
	fmt.Println("5 untuk ubah harga sewa")
	fmt.Print("Pilihan : ")
	fmt.Scan(&userChoice)
	fmt.Println("")
	return userChoice
}

func UbahDataNamaTempat(dataTempat *tabTempat, indexTempat int) {
	var namaBaru string
	fmt.Print("Masukan nama baru : ")
	fmt.Scan(&namaBaru)
	dataTempat[indexTempat].namaTempat = namaBaru
	fmt.Println("Nama berhasil diubah.")
	fmt.Println("")
}

func UbahDataLokasiTempat(dataTempat *tabTempat, indexTempat int) {
	var lokasiBaru string
	fmt.Print("Masukan lokasi baru : ")
	fmt.Scan(&lokasiBaru)
	dataTempat[indexTempat].lokasiTempat = lokasiBaru
	fmt.Println("lokasi berhasil diubah.")
	fmt.Println("")
}

func UbahDataKapasitasTempat(dataTempat *tabTempat, indexTempat int) {
	var kapasitasBaru int
	fmt.Print("Masukan kapasitas baru : ")
	fmt.Scan(&kapasitasBaru)
	dataTempat[indexTempat].kapasitasMaksimum = kapasitasBaru
	fmt.Println("kapasitas berhasil diubah.")
	fmt.Println("")
}

func UbahDataFasilitasTempat(dataTempat *tabTempat, indexTempat int) {
	var userChoice int = -1
	var fasilitasBaru string
	fmt.Println("Data fasilitas :")
	for i := 0; i < len(dataTempat[indexTempat].fasilitas); i++ {
		fmt.Printf("  %d. %s\n", i+1, dataTempat[indexTempat].fasilitas[i])
	}
	for userChoice <= 0 || userChoice > len(dataTempat[indexTempat].fasilitas) {
		fmt.Print("Pilih no fasilitas yang ingin anda ubah : ")
		fmt.Scan(&userChoice)
		if userChoice > 0 && userChoice <= len(dataTempat[indexTempat].fasilitas) {
			fmt.Print("masukan fasilitas baru : ")
			fmt.Scan(&fasilitasBaru)
		}
	}

	dataTempat[indexTempat].fasilitas[userChoice-1] = fasilitasBaru
	fmt.Println("fasilitas berhasil diubah.")
	fmt.Println("")

}

func UbahDataHargaTempat(dataTempat *tabTempat, indexTempat int) {
	var hargaBaru int
	fmt.Print("Masukan harga baru : ")
	fmt.Scan(&hargaBaru)
	dataTempat[indexTempat].hargaSewa = hargaBaru
	fmt.Println("harga sewa berhasil diubah.")
	fmt.Println("")
}

func cetakDataTempat(dataTempat tabTempat, indexTempat int) {
	fmt.Println("Nama Tempat :", dataTempat[indexTempat].namaTempat)
	fmt.Println("Lokasi Tempat :", dataTempat[indexTempat].lokasiTempat)
	fmt.Println("Kapasitas Maksimum :", dataTempat[indexTempat].kapasitasMaksimum)
	fmt.Println("Fasilitas :")
	for i := 0; i < len(dataTempat[indexTempat].fasilitas); i++ {
		fmt.Printf("  %d. %s\n", i+1, dataTempat[indexTempat].fasilitas[i])
	}
	//Loop untuk print setiap fasilitas tempat (list)
	fmt.Println("Harga sewa :", dataTempat[indexTempat].hargaSewa)
	fmt.Println("Riwayat sewa : -")
	for i := 0; i < len(dataTempat[indexTempat].riwayatSewa); i++ {
		fmt.Printf("  %d. %s, %d, %d - %d", i+1, dataTempat[indexTempat].riwayatSewa[i].namaPenyewa, dataTempat[indexTempat].riwayatSewa[i].tanggal, dataTempat[indexTempat].riwayatSewa[i].jamMulai, dataTempat[indexTempat].riwayatSewa[i].jamSelesai)
	}
	fmt.Println("")
}

func fiturTampilkanData(dataTempat tabTempat, nDataTempat int) {
	for i := 0; i < nDataTempat; i++ {
		fmt.Println("TEMPAT", i+1)
		cetakDataTempat(dataTempat, i)
		fmt.Println()
	}
}

func fiturHapusDataTempat(dataTempat *tabTempat, nDataTempat *int) {
	var indexTempat int
	var userChoice string
	fmt.Println("Silahkan cari tempat terlebih dahulu")
	indexTempat = cariTempat(*dataTempat, *nDataTempat)
	cetakDataTempat(*dataTempat, indexTempat)
	fmt.Print("Lanjut menghapus data diatas (y/n) ? ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		hapusDataTempat(dataTempat, nDataTempat, indexTempat)
	}
	fmt.Println()
	fiturTampilkanData(*dataTempat, *nDataTempat)
}

func hapusDataTempat(dataTempat *tabTempat, nDataTempat *int, indexTempat int) {
	for i := indexTempat + 1; i < *nDataTempat; i++ {
		dataTempat[i-1] = dataTempat[i]
	}
	*nDataTempat -= 1
	fmt.Println("Data berhasil dihapus.")
	fmt.Println("")
}

func fiturUrutDataTempat(dataTempat *tabTempat, nDataTempat int) {
	var userChoice int
	fmt.Println("Fitur Tampilkan Data Terurut.")
	userChoice = menuUrutDataTempat()

	switch userChoice {
	case 1:
		urutDenganNama(dataTempat, nDataTempat)
	case 2:
		urutDenganLokasi(dataTempat, nDataTempat)
	case 3:
		urutDengaKapasitas(dataTempat, nDataTempat)
	case 4:
		urutDenganHargaSewa(dataTempat, nDataTempat)
	case 5:
	}

}

func menuUrutDataTempat() int {
	var userChoice int
	for userChoice < 1 || userChoice > 5 {
		fmt.Println("Pilih metode urut data. Ketik : ")
		fmt.Println("1 untuk mengurutkan berdasarkan nama")
		fmt.Println("2 untuk mengurutkan berdasarkan lokasi")
		fmt.Println("3 untuk mengurutkan berdasarkan kapasitas")
		fmt.Println("4 untuk mengurutkan berdasarkan hargasewa")
		fmt.Print("Pilihan : ")
		fmt.Scan(&userChoice)
		fmt.Println("")
		if userChoice >= 1 && userChoice <= 4 {
			return userChoice
		} else {
			fmt.Println("Pilihan tidak tersedia, silahkan ulangi.")
		}
	}

	return userChoice
}

func urutDenganNama(dataTempat *tabTempat, nDataTempat int) {
	var i, idx, pass int
	var temp iTempat

	for pass = 0; pass < nDataTempat-1; pass++ {
		idx = pass

		for i = pass + 1; i < nDataTempat; i++ {
			if dataTempat[i].namaTempat < dataTempat[idx].namaTempat {
				idx = i
			}
		}
		temp = dataTempat[idx]
		dataTempat[idx] = dataTempat[pass]
		dataTempat[pass] = temp
	}
}

func urutDenganLokasi(dataTempat *tabTempat, nDataTempat int) {
	var i, pass int
	var temp iTempat

	for pass = 1; pass <= nDataTempat-1; pass++ {
		temp = dataTempat[pass]
		i = pass

		for i > 0 && dataTempat[i-1].lokasiTempat > temp.lokasiTempat {
			dataTempat[i] = dataTempat[i-1]
			i--
		}
		dataTempat[i] = temp
	}
}

func urutDengaKapasitas(dataTempat *tabTempat, nDataTempat int) {
	var i, idx, pass int
	var temp iTempat

	for pass = 0; pass < nDataTempat-1; pass++ {
		idx = pass

		for i = pass + 1; i < nDataTempat; i++ {
			if dataTempat[i].kapasitasMaksimum < dataTempat[idx].kapasitasMaksimum {
				idx = i
			}
		}
		temp = dataTempat[idx]
		dataTempat[idx] = dataTempat[pass]
		dataTempat[pass] = temp
	}
}

func urutDenganHargaSewa(dataTempat *tabTempat, nDataTempat int) {
	var i, pass int
	var temp iTempat

	for pass = 1; pass <= nDataTempat-1; pass++ {
		temp = dataTempat[pass]
		i = pass

		for i > 0 && dataTempat[i-1].hargaSewa > temp.hargaSewa {
			dataTempat[i] = dataTempat[i-1]
			i--
		}
		dataTempat[i] = temp
	}

}
