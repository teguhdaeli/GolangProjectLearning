/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//named return values --> bisa memberikan nama variabel pada return values

package main

import "fmt"

func getDataMhs() (firstName, MiddleName string, umur int){
	firstName = "Teguh"
	MiddleName = "Daeli"
	umur = 22
	return //<-- cara pertama
	//return firstName, MiddleName, umur <-- bisa cara lain atau cara kedua
}

func getPerkalian() (nn1, nn2, hasil2 int){
	nn1 = 53
	nn2 = 39
	hasil2 = 0

	return nn1, nn2, hasil2
}
func main(){
	//cara memanggil atau munculkan outputnya
	//buat lagi variabelnya
	firstName, MiddleName, umur := getDataMhs() //<-- cara pertama untuk variable
	fmt.Println(firstName, MiddleName, umur)
	//
	//return get perkalian
	a,b,c := getPerkalian() //bisa cara kedua untuk menyingkat nama variabel di function
	c = a * b
	fmt.Println(a," x ",b, "=", c)

}