/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

/*
Struct >> template data untuk menggabungkan nol atau lebih dari tipe data lainnya dalam satu kesatuan
Struct >> kumpulan dari field atau banyak tipe data
*/


package main

import "fmt"
type Mahasiswa struct{
	nim, nama string
	umur, tinggi int
}

func main(){
	//cara pertama isi struct
	var mhs Mahasiswa
	mhs.nim = "182110960"
	mhs.nama = "Teguh Christianto Daeli"
	mhs.umur = 22
	mhs.tinggi = 180

	//Pemanggilan
	fmt.Println(mhs)
	//pemanggilan satu persatu
	fmt.Println(mhs.nama)
	fmt.Println(mhs.nim)

	//cara kedua isi struct
	mhsMikroskil := Mahasiswa{
		nim: "18211160",
		nama: "Joko Jiwono",
		umur: 25,
		tinggi: 192,
	}
	fmt.Println(mhsMikroskil)
}