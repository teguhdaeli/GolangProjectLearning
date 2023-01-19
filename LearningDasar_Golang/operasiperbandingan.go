/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/
//operasi perbanding untuk membandingkan dua buah data

package main

import "fmt"

func main(){
	//>, <, >=, <=, ==, !=

	var name1= "Teguh"
	var name2= "Daeli"

	var result bool = name1 == name2

	fmt.Println(result)

	//
	var NilaiUtama = 100
	var NilaKedua = 300

	var result2 = NilaiUtama < NilaKedua
	fmt.Println(result2)

}