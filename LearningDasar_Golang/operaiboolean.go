/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//operasi boolean
//&&, ||, !

package main

import "fmt"

func main(){
	// mengecek kelulusan dengan operasi boolean
	var NilaiAkhir = 80
	var NilaiAbsen = 80

	var lulusUjian = NilaiAkhir < 70
	var lulusAbsen = NilaiAbsen < 70

	var hasilLulus bool = lulusUjian && lulusAbsen
	fmt.Println(hasilLulus)
}

