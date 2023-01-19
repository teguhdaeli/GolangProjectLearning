/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//Array

package main

import "fmt"

func main(){
	//membuat array dengan jumlah yang sudahh ditentukan
	var potonganNama[3] string
	potonganNama[0] = "Teguh"
	potonganNama[1] = "Christianto"
	potonganNama[2] = "Daeli"
	//
	var umur[2] int
	umur[0] = 20
	umur[1] = 21
	//
	//Membuat array isi data sekaligus
	var values =[3] int {90, 91, 92}
	//cara memanggil isi array 
	fmt.Println(potonganNama, umur[1])
	fmt.Println(values[2], values)
	//
	//function array, len(array) untuk melihat panjang data, array[index], array[index] = value
	fmt.Println("Panjangan data array : ", len(potonganNama))
	fmt.Println("Panjangan data array : ", len(potonganNama))
	fmt.Println("Panjangan data array : ", len(potonganNama))

}