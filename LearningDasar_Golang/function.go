/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

/*
Function --> sebuah blok kode yang dibuat, agar bisa di jalankan berulang ulang

*/

package main

import "fmt"

//Pembuatan function
func person(){
	person := make(map[string]string)
	person["Name"] = "Teguh"
	person["Umur"] = "22"
	fmt.Println(person)
}

func main(){
	person()

}