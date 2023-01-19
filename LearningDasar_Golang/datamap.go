/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/
//Data map --> kumpulan data, dimana kita bisa menentukan jenis data index yang kita gunakan
// terdiri dari key - value(kata kunci - nilai)

package main 

import "fmt"

func main(){
	// membuat map
	 person := map[string]string{
		"Nama"	: "Teguh",
		"Umur"	: "22",
	}

	fmt.Println(person)
	//
	//function map
	fmt.Println(len(person))

	fmt.Println(person["Nama"])

	//membuat map baru
	var book map[string]string = make(map[string]string)
	book["Judul"] = "Belajar Go"
	book["Author"] = "Teguh Daeli"
	book["Tahun"] = "2022"
	fmt.Println(book)
	delete(book,"Tahun")
	fmt.Println(book)



}