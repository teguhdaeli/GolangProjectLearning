/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//Type deklarasi membuat ulang tipe data baru yang sudah ada

package main

import "fmt"

func main(){
	type NomorHp string
	
	var nomorTeguh NomorHp = "082162798664"
	fmt.Println(NomorHp, nomorTeguh)


}