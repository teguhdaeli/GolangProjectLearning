/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

/*
Function as parameter dapat kita passing ke function lain

*/

package main

import "fmt"

func sayHellowithFilter(name string, filter func(string)string){ // cara pertama
	fmt.Println("Hello", filter(name))
}

//cara kedua bisa menggunakan type declarations
type Filter func(string)string
func cekFamily(nameFamily string, filter Filter){
	fmt.Println("Hallo Family : ", filter(nameFamily))
}

func filterFamily(nameFamily string)string{
	if nameFamily == "Teguh"{
		return nameFamily
	}else{
		return "Tidak diketahui!"
	}
}
//
func spamFilter(name string)string{
	if name == "anjing"{
		return "..."
	} else{
		return name
	}
}
func main(){
	sayHellowithFilter("Teguh", spamFilter)
	sayHellowithFilter("anjing", spamFilter)
	cekFamily("as", filterFamily)
}