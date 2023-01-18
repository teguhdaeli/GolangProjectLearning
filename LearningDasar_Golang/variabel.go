/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

package main

import "fmt"

func main() {
	//create variabel
	var nameBoy, Status, nameGirl string
	nameBoy = "Teguh Christianto Daeli"
	Status = "Suami_Istri"
	nameGirl = "Nita Kwok bauk"

	// variabel tanpa penyebutan jenis data nya
	var umurPria =22
	var umurWanita=22

	agamaBerdua := "Kristen"
	Country := 20
	
	//Variable banyak data
	var(
		firsName = "Teguh"
		lastName = "Daeli"
		Umur = 22
	)

	fmt.Println("Nama Pasangan Pria	    : ", nameBoy)
	fmt.Println("Status Hubungan	    : ", Status)
	fmt.Println("Nama Pasangan Wanita	: ", nameGirl)
	fmt.Println(umurPria, umurWanita)
	fmt.Println(agamaBerdua, Country)
	fmt.Println(firsName, lastName,Umur)
}
