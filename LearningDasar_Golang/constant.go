/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/
package main

import "fmt"
//Pemakaian Constant variabel yang tidak bisa diubah lagi isi datanya, dan hanya bisa dipakai sekali
func main(){
	const fisrtName string= "Teguh"
	const lastName = "Daeli"
	const value = 22
	const tinggi = 169

	//constant multiple
	const(
		agama string="Kristen"
		bahasa = "Indonesia"
		lamaTinggal = 23

	)

	fmt.Println(fisrtName)
	fmt.Println(tinggi)
	fmt.Println(lastName)
	fmt.Println(value)
	fmt.Println(agama)
	fmt.Println(bahasa)
	fmt.Println(lamaTinggal)
}