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
	totalBiaya  int
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
		namaTempat:        "gwk",
		lokasiTempat:      "bali",
		kapasitasMaksimum: 100,
		fasilitas:         []string{"taman", "kantin"},
		hargaSewa:         1000000,
		riwayatSewa:       []waktuSewa{},
	}
	dataTempat[1] = iTempat{
		namaTempat:        "gbk",
		lokasiTempat:      "jakarta",
		kapasitasMaksimum: 200,
		fasilitas:         []string{"lapangan bola", "lapangan basket"},
		hargaSewa:         2000000,
		riwayatSewa:       []waktuSewa{},
	}
	dataTempat[2] = iTempat{
		namaTempat:        "borobudur",
		lokasiTempat:      "magelang",
		kapasitasMaksimum: 400,
		fasilitas:         []string{"kantin", "taman"},
		hargaSewa:         1500000,
		riwayatSewa:       []waktuSewa{},
	}
	*nDataTempat = 3
}

func main() {
	var dataTempat tabTempat
	var nDataTempat int
	var dataUser tabUserAcc
	var nDataUser int

	//ADD DUMMY ACCOUNT FOR TESTING PURPOSE ONLY
	testingPurpose(&dataTempat, &dataUser, &nDataTempat, &nDataUser)

	var currentUserClass string // user class dari akun yang sedang dipakai "manajer" OR "pelanggan"
	var currentIndexAcc int     // index dari akun yang sedang dipakai

	mainMenu(&dataTempat, &nDataTempat, &dataUser, &nDataUser, &currentUserClass, &currentIndexAcc)
}

// PUBLIC's FUNCTION START HERE

func mainMenu(dataTempat *tabTempat, nDataTempat *int, dataUser *tabUserAcc, nDataUser *int, currentUserClass *string, currentIndexAcc *int) {
	/*
		 	IS : 	-
			FS : 	Mengembalikan string "login" atau "register" atau Keluar program sesuai dengan pilihan user ketika mengisi variable userChoice
	*/
	var userChoice int = 0
	for userChoice != 3 {
		fmt.Println("Menu utama, silahkan masukan angka :")
		fmt.Println("1 untuk login")
		fmt.Println("2 untuk daftar")
		fmt.Println("3 untuk keluar program")
		fmt.Print("pilihan : ")
		fmt.Scan(&userChoice)
		fmt.Println("")
		if userChoice == 1 {
			login(*dataUser, *nDataUser, currentUserClass, currentIndexAcc)
		} else if userChoice == 2 {
			register(dataUser, nDataUser, currentUserClass, currentIndexAcc)
		} else if userChoice == 3 {
			*currentIndexAcc = -1
			*currentUserClass = "0"
			fmt.Print("Terimakasih Telah Menggunakan Program Sewa Tempat")
		} else {
			fmt.Println("\npilihan tidak tersedia, silahkan ulangi.")
		}

		if userChoice == 1 || userChoice == 2 {
			if *currentUserClass == "pelanggan" {
				menuPelanggan(dataTempat, *nDataTempat, currentIndexAcc, currentUserClass)
			} else if *currentUserClass == "manajer" {
				menuManajer(dataTempat, nDataTempat, currentIndexAcc, currentUserClass)
			}
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
			fmt.Println("Akun berhasil login")
			break
		}
	}
	if found == false {
		fmt.Println("\n\nAkun tidak ditemukan. Silahkan ulangi.")
		login(dataUser, nDataUser, currentUserClass, currentIndexAcc)
	}
	fmt.Println("")
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
		fmt.Println("")
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

func cariTempat(dataTempat *tabTempat, nDataTempat int) int {
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
		fmt.Println("4 untuk mencari dengan harga sewa")
		fmt.Print("masukan pilihan anda : ")
		fmt.Scan(&userChoice)
		fmt.Println("")

		switch userChoice { // ganti if else biasa
		case 1:
			indexTempat = cariDenganNama(*dataTempat, nDataTempat)
		case 2:
			urutDenganLokasi(dataTempat, nDataTempat)
			indexTempat = cariDenganLokasi(*dataTempat, nDataTempat)
		case 3:
			indexTempat = cariDenganKapasitas(*dataTempat, nDataTempat)
		case 4:
			urutDenganHargaSewa(dataTempat, nDataTempat)
			indexTempat = cariDenganHarga(*dataTempat, nDataTempat)
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
	/*
			IS :	variabel dataTempat dan nDataTempat yang menyimpan datar tempat dan jumlah tempat
			FS : 	mengembalikan index dari nama yang cocok denga yang user cari, mengembalikan -1 jika tidak ada yg cocok
   				(menggunakan Sequential Search)
	*/
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
	/*
			IS :	variabel dataTempat dan nDataTempat yang menyimpan datar tempat dan jumlah tempat
			FS : 	mengembalikan index dari nama yang cocok denga yang user cari, mengembalikan -1 jika tidak ada yg cocok
   				(menggunakan Binary Search)
	*/
	var lokasiYangDicari string
	fmt.Println("Pencarian.")
	fmt.Print("Masukan lokasi tempat : ")
	fmt.Scan(&lokasiYangDicari)
	fmt.Println("")

	L, R := 0, nDataTempat-1
	for L <= R {
		mid := (L + R) / 2
		if dataTempat[mid].lokasiTempat < lokasiYangDicari {
			L = mid + 1
		} else if dataTempat[mid].lokasiTempat > lokasiYangDicari {
			R = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func cariDenganKapasitas(dataTempat tabTempat, nDataTempat int) int {
	/*
			IS :	variabel dataTempat dan nDataTempat yang menyimpan datar tempat dan jumlah tempat
			FS : 	mengembalikan index dari nama yang cocok denga yang user cari, mengembalikan -1 jika tidak ada yg cocok
   				(menggunakan Sequential Search)
	*/
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

func cariDenganHarga(dataTempat tabTempat, nDataTempat int) int {
	/*
			IS :	variabel dataTempat dan nDataTempat yang menyimpan datar tempat dan jumlah tempat
			FS : 	mengembalikan index dari nama yang cocok denga yang user cari, mengembalikan -1 jika tidak ada yg cocok
   				(menggunakan Binary Search)
	*/
	var hargaYangDicari int
	fmt.Print("Pencarian.")
	fmt.Print("Masukan harga tempat : ")
	fmt.Scan(&hargaYangDicari)
	fmt.Println("")
	L, R := 0, nDataTempat-1
	for L <= R {
		mid := (L + R) / 2
		if dataTempat[mid].hargaSewa < hargaYangDicari {
			L = mid + 1
		} else if dataTempat[mid].hargaSewa > hargaYangDicari {
			R = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func menuPelanggan(dataTempat *tabTempat, nDataTempat int, currentIndexAcc *int, currentUserClass *string) {
	/*

		IS : variabel data tempat,nDataTempat,currentindexAcc dan currentuserClass menyimpan data tempat, jumlah data tempat, index akun user saat ini dan class user
		FS : Memanggil fitur sewa tempat dan menginput data dari user kemudian memanggil isAvailabele untuk mengecek ketersediaan tempat,
		ketika kondisi true maka riwayat penyewaan user dari input sewa tempat akan ditampilkan

	*/
	var userChoice int = 0

	for userChoice != 2 {
		fmt.Println("Menu utama pelanggan, Silahkan ketik")
		fmt.Println("1 untuk sewa tempat")
		fmt.Println("2 untuk keluar")
		fmt.Print("Pilihan : ")
		fmt.Scan(&userChoice)
		fmt.Println("")

		if userChoice == 1 {
			fiturSewaTempat(dataTempat, nDataTempat)
		} else if userChoice == 2 {
			Logout(currentIndexAcc, currentUserClass)
			fmt.Println("Berhasil keluar")
			fmt.Println("")
		} else {
			fmt.Println("Pilihan tidak tersedia, silahkan ulangi.")
			//menuPelanggan(dataTempat, nDataTempat, currentIndexAcc, currentUserClass) --> jika dipanggil dalam loop, tidak perlu dibuat rekrusif
		}

	}
}

func fiturSewaTempat(dataTempat *tabTempat, nDataTempat int) {
	var currentIndexTempat int
	currentIndexTempat = cariTempat(dataTempat, nDataTempat)

	if currentIndexTempat == -1 {
		fmt.Println("Tempat tidak tersedia.")
		return
	}
	// Proses Booking Venue
	var booking waktuSewa
	var durasi int

	fmt.Print("Masukkan nama penyewa: ")
	fmt.Scan(&booking.namaPenyewa) // nama penyewa diambil dari username akun yang sudah terlogin
	fmt.Print("Masukkan tanggal sewa (format: YYYYMMDD): ")
	fmt.Scan(&booking.tanggal)
	fmt.Print("Masukkan jam mulai (format 24 jam): ")
	fmt.Scan(&booking.jamMulai)
	fmt.Print("Masukkan durasi sewa (dalam jam): ")
	fmt.Scan(&durasi) // input nya durasi
	booking.jamSelesai = booking.jamMulai + durasi
	booking.totalBiaya = durasi * dataTempat[currentIndexTempat].hargaSewa

	// Cek apakah tempat tersedia di waktu yang diminta
	if isAvailable(dataTempat[currentIndexTempat], booking) {
		dataTempat[currentIndexTempat].riwayatSewa = append(dataTempat[currentIndexTempat].riwayatSewa, booking)
		fmt.Println("Tempat berhasil disewa.")
		fmt.Println("")
		tampilkanRiwayatSewa(dataTempat[currentIndexTempat])
	} else {
		fmt.Println("Tempat tidak tersedia pada waktu yang diminta.")
		fmt.Println("")
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
		fmt.Println("")
		fmt.Printf("Nama Penyewa: %s\n", sewa.namaPenyewa)
		fmt.Printf("Tanggal : %d\n", sewa.tanggal)
		fmt.Printf("Jam mulai : %d\n", sewa.jamMulai)
		fmt.Printf("Jam Selesai : %d\n", sewa.jamSelesai)
		fmt.Printf("Total biaya : %d\n", sewa.totalBiaya)
		fmt.Println("")
	}
}

func menuManajer(dataTempat *tabTempat, nDataTempat *int, currentIndexAcc *int, currentUserClass *string) {
	/*
		IS : dataTempat berisi data tempat, nDataTempat berisi jumlah tempat, currentIndexAcc berisi index user saat ini
		currentUserClass berisi class dari user
		FS : opsi selanjutnya bergantung pada jawabab yang dipilih user antara menampilkan data terurut, merubah data tempat,
		menghapus data tempat atau logout dari menu
	*/
	var userChoice int = 0
	for userChoice != 4 {
		fmt.Println("Menu utama manajer, Silahkan ketik")
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
			fmt.Println("Berhasil keluar")
			fmt.Println("")
		} else {
			fmt.Println("Pilian tidak tersedia, silahkan ulangi.")
		}
	}

}

func fiturUbahDataTempat(dataTempat *tabTempat, nDataTempat int) {
	/*
		IS : dataTempat berisi data tempat dan nDataTempat berisi jumlah data tempat
		FS : mengubah data tempat pilihan user, dengan data baru yang diinput user
	*/
	var indexTempat, userChoice int
	fmt.Println("Silahkan cari tempat terlebih dahulu")
	indexTempat = cariTempat(dataTempat, nDataTempat)

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
	/*
		IS : dataTempat berisi data tempat dan indexTemat berisi index dari tempat yang ingin diubah
		FS : menampilkan menu data yang ingin diubah, mengembalikan integer sesuai dengan input dari user
	*/
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
	/*
		IS : dataTempat berisi data tempat dan indexTemat berisi index dari tempat yang ingin diubah
		FS : mengubah nilai dari namaTempat dari iTempat array ke indexTempat
	*/
	var namaBaru string
	fmt.Print("Masukan nama baru : ")
	fmt.Scan(&namaBaru)
	dataTempat[indexTempat].namaTempat = namaBaru
	fmt.Println("Nama berhasil diubah.")
	fmt.Println("")
}

func UbahDataLokasiTempat(dataTempat *tabTempat, indexTempat int) {
	/*
		IS : dataTempat berisi data tempat dan indexTemat berisi index dari tempat yang ingin diubah
		FS : mengubah nilai dari lokasiTempat dari iTempat array ke indexTempat
	*/
	var lokasiBaru string
	fmt.Print("Masukan lokasi baru : ")
	fmt.Scan(&lokasiBaru)
	dataTempat[indexTempat].lokasiTempat = lokasiBaru
	fmt.Println("lokasi berhasil diubah.")
	fmt.Println("")
}

func UbahDataKapasitasTempat(dataTempat *tabTempat, indexTempat int) {
	/*
		IS : dataTempat berisi data tempat dan indexTemat berisi index dari tempat yang ingin diubah
		FS : mengubah nilai dari kapasitasMaksimum dari iTempat array ke indexTempat
	*/
	var kapasitasBaru int
	fmt.Print("Masukan kapasitas baru : ")
	fmt.Scan(&kapasitasBaru)
	dataTempat[indexTempat].kapasitasMaksimum = kapasitasBaru
	fmt.Println("kapasitas berhasil diubah.")
	fmt.Println("")
}

func UbahDataFasilitasTempat(dataTempat *tabTempat, indexTempat int) {
	/*
		IS : dataTempat berisi data tempat dan indexTemat berisi index dari tempat yang ingin diubah
		FS : mengubah nilai dari fasilitas dari iTempat array ke indexTempat
	*/
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
	/*
		IS : dataTempat berisi data tempat dan indexTemat berisi index dari tempat yang ingin diubah
		FS : mengubah nilai dari hargaSewa dari iTempat array ke indexTempat
	*/
	var hargaBaru int
	fmt.Print("Masukan harga baru : ")
	fmt.Scan(&hargaBaru)
	dataTempat[indexTempat].hargaSewa = hargaBaru
	fmt.Println("harga sewa berhasil diubah.")
	fmt.Println("")
}

func cetakDataTempat(dataTempat tabTempat, indexTempat int) {
	/*
		IS : dataTempat berisi data tempat dan indexTemat berisi index dari tempat yang ingin diubah
		FS : menampilkan data dari array dataTempat dengan index ke indexTempat
	*/
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
	/*
		IS : dataTempat berisi data tempat dan nDataTempat yang berisi nilai banyaknya data dalam dataTempat
		FS : menampilkan semua data di dataTempat dengan memanggil cetakData dengan loop
	*/
	for i := 0; i < nDataTempat; i++ {
		fmt.Println("TEMPAT", i+1)
		cetakDataTempat(dataTempat, i)
		fmt.Println()
	}
}

func fiturHapusDataTempat(dataTempat *tabTempat, nDataTempat *int) {
	/*
		IS : dataTempat berisi data tempat dan nDataTempat yang berisi nilai banyaknya data dalam dataTempat
		FS : menghapus data denga index sesuai pilihan user
	*/
	var indexTempat int
	var userChoice string
	fmt.Println("Silahkan cari tempat terlebih dahulu")
	indexTempat = cariTempat(dataTempat, *nDataTempat)
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
	/*
		IS : dataTempat berisi data tempat dan nDataTempat yang berisi nilai banyaknya data dalam dataTempat, indexTempat yang berisi index dari tempat yang dipilih user
		FS : memajukan index dari dataTempat mulai dari index indexTempat + 1, mengurangi nDataTempat dengan 1
	*/
	for i := indexTempat + 1; i < *nDataTempat; i++ {
		dataTempat[i-1] = dataTempat[i]
	}
	*nDataTempat -= 1
	fmt.Println("Data berhasil dihapus.")
	fmt.Println("")
}

func fiturUrutDataTempat(dataTempat *tabTempat, nDataTempat int) {
	/*
		IS : dataTempat berisi data tempat dan nDataTempat berisi jumlah data tempat
		FS : dataTempat terurut sesuai urutan pilihan user
	*/
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
	/*
		IS : -
		FS : mengembalikan userChoice yang berisi nilai masukan dari user sesuai pilihan metode urut
	*/
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
	/*
		IS : 	dataTempat berisi data tempat dan nDataTempat berisi jumlah data tempat
		FS : 	dataTempat terurut Acending berdasarkan namaTempat
  			(menggunakan Selection sort)
	*/
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
	/*
		IS : 	dataTempat berisi data tempat dan nDataTempat berisi jumlah data tempat
		FS : 	dataTempat terurut Acending berdasarkan lokasiTempat
  			(menggunakan Insertion sort)
	*/
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
	/*
		IS : 	dataTempat berisi data tempat dan nDataTempat berisi jumlah data tempat
		FS : 	dataTempat terurut Acending berdasarkan kapasitasMaksimum
  			(menggunakan Selection sort)
	*/
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
	/*
		IS : 	dataTempat berisi data tempat dan nDataTempat berisi jumlah data tempat
		FS : 	dataTempat terurut Acending berdasarkan hargaSewa
  			(menggunakan Insertion sort)
	*/
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
