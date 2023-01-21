 /*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/
/*
Struct method dimana dia bisa digunakan sebaga parameter untuk function
Method adalah function
*/
package main

import "fmt"

type Mahasiswa struct{
	nim, nama string
	umur, tinggi int
}

func(mhs Mahasiswa) sayMahasiswa(){
	fmt.Println("Hallo, ", mhs.nama)
}
func main(){
	//cara pertama isi struct
	var mhs Mahasiswa
	mhs.nim = "182110960"
	mhs.nama = "Teguh Christianto Daeli"
	mhs.umur = 22
	mhs.tinggi = 180
	mhs.sayMahasiswa()
}
