/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//Cara konversi tipe data di golang

package main

import "fmt"

func main(){
	var Nilai32 int32 = 129
	var nilai64 int64 = int64(Nilai32)

	fmt.Println(Nilai32)
	fmt.Println(nilai64)
}