package main

import "fmt"

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

func main() {
	var dataTempat tabTempat
	var nDataTempat int
	var dataUser tabUserAcc
	var nDataUser int

	var currentUserClass string // user class dari akun yang sedang dipakai "manajer" OR "pelanggan"
	var currentIndexAcc int     // index dari akun yang sedang dipakai

	testingPurpose(&dataTempat, &dataUser, &nDataTempat, &nDataUser)

	if loginOrRegister() == "login" {
		login(dataUser, nDataUser, &currentUserClass, &currentIndexAcc)
	} else {
		register(&dataUser, &nDataUser, &currentUserClass, &currentIndexAcc)
	}

	//Test the connected account
	ConnectedAccountTest(dataUser, nDataUser, currentUserClass, currentIndexAcc)

}

func loginOrRegister() string {
	var userChoice int
	fmt.Println("Masukan angka :")
	fmt.Println("1 untuk login")
	fmt.Println("2 untuk daftar")
	fmt.Print("pilihan : ")
	fmt.Scan(&userChoice)
	fmt.Println("")

	if userChoice == 1 {
		return "login"
	} else if userChoice == 2 {
		return "register"
	} else {
		fmt.Println("\npilihan tidak tersedia, silahkan ulangi.")
		return loginOrRegister()
	}
}

func login(dataUser tabUserAcc, ndataUser int, currentUserClass *string, currentIndexAcc *int) {
	var username, password string
	var found bool = false
	fmt.Println("Login. Silahkan masukan data.")
	fmt.Print("username : ")
	fmt.Scan(&username)
	fmt.Print("password : ")
	fmt.Scan(&password)

	for i := 0; i < ndataUser; i++ {
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
		login(dataUser, ndataUser, currentUserClass, currentIndexAcc)
	}
}

func register(dataUser *tabUserAcc, ndataUser *int, currentUserClass *string, currentIndexAcc *int) {
	var username, password string
	var found bool = false
	fmt.Println("Register. Silahkan masukan data.")
	fmt.Print("username : ")
	fmt.Scan(&username)
	fmt.Print("password : ")
	fmt.Scan(&password)

	for i := 0; i < *ndataUser; i++ {
		if dataUser[i].userName == username {
			found = true
			break
		}
	}

	if found == true {
		fmt.Println("\n\nUsername sudah dipakai. Silahkan ulangi.")
		register(dataUser, ndataUser, currentUserClass, currentIndexAcc)
	} else {
		newUser := userAcc{
			userName:  username,
			password:  password,
			userClass: "pelanggan",
		}
		dataUser[*ndataUser] = newUser
		*ndataUser++
		*currentUserClass = newUser.userClass
		*currentIndexAcc = *ndataUser - 1
		fmt.Println("Akun berhasil register")
	}
}

func Logout(currentIndexAcc *int, currentUserClass *string) {
	*currentIndexAcc = -1
	*currentUserClass = "0"
}

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

func ConnectedAccountTest(dataUser tabUserAcc, nDataUser int, currentUserClass string, currentIndexAcc int) {
	fmt.Println("currentIndexAcc : ", currentIndexAcc)
	fmt.Println("dataUser[currentIndexAcc].userName :", dataUser[currentIndexAcc].userName)
	fmt.Println("dataUser[currentIndexAcc].password :", dataUser[currentIndexAcc].password)
	fmt.Print("dataUser[currentIndexAcc].userClass :", dataUser[currentIndexAcc].userClass)
	fmt.Print("currentUserClass : ", currentUserClass)
}

// FITUR PELANGGAN // FITUR PELANGGAN // FITUR PELANGGAN // FITUR PELANGGAN // FITUR PELANGGAN // FITUR PELANGGAN // FITUR PELANGGAN //

func cariTempat(dataTempat tabTempat, nDataTempat int) int {
	var userChoice int
	var indexTempat int
	fmt.Println("pilih metode pencarian. ketik : ")
	fmt.Println("1 untuk mencari dengan nama")
	fmt.Println("2 untuk mencari dengan lokasi")
	fmt.Println("3 untuk mencari dengan kapasitas")
	fmt.Println("masukan pilihan anda : ")
	fmt.Scan(&userChoice)

	switch userChoice {
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
		return cariTempat(dataTempat, nDataTempat)
	} else {
		return indexTempat
	}
}

func cariDenganNama(dataTempat tabTempat, nDataTempat int) int {
	var namaYangDicari string
	fmt.Print("Masukan : ")
	fmt.Scan(&namaYangDicari)

	// Looping mencari data berdasarkan data tempat
	for i := 0; i < nDataTempat; i++ {
		if namaYangDicari == dataTempat[i].namaTempat {
			return i
		}
	}
	return -1
}

func cariDenganLokasi(dataTempat tabTempat, nDataTempat int) int {
	var lokasiYangDicari string
	fmt.Print("Masukan : ")
	fmt.Scan(&lokasiYangDicari)
	for i := 0; i < nDataTempat; i++ {
		if lokasiYangDicari == dataTempat[i].lokasiTempat {
			return i
		}
	}
	return -1
}

func cariDenganKapasitas(dataTempat tabTempat, nDataTempat int) int {
	var kapasitasYangDicari int
	fmt.Print("Masukan : ")
	fmt.Scan(&kapasitasYangDicari)
	for i := 0; i < nDataTempat; i++ {
		if kapasitasYangDicari == dataTempat[i].kapasitasMaksimum {
			return i
		}
	}
	return -1
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
	fmt.Scan(&booking.namaPenyewa)
	fmt.Println("Masukkan tanggal sewa (format: YYYYMMDD): ")
	fmt.Scan(&booking.tanggal)
	fmt.Println("Masukkan jam mulai (format 24 jam): ")
	fmt.Scan(&booking.jamMulai)
	fmt.Println("Masukkan jam selesai (format 24 jam): ")
	fmt.Scan(&booking.jamSelesai)

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

//Bikin Prossedure sewa tempat
