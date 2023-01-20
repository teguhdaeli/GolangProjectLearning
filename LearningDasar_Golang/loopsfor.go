/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//For loop

package main

import "fmt"

func main(){
	/**for
	i := 0
	for i <= 10{
		fmt.Println("Perulangan Ke : ", i)
		i++
	}

	//for dengan cara lain atua for in statement
	for i := 0; i<10; i++{
		fmt.Println("Perulangan Dengan : ", i)
	}
	**/

	slice := []string{"Teguh","Nita", "Wulan", "Delvin"}
	for i := 0; i < len(slice); i++{
		fmt.Println("Name ke :", i, " -->", slice[i])
	}
	//for range

	for key, value := range slice{ //<-- jika tidak ingin ada nomor key atau index bisa gunakanan __ atau underscore
		fmt.Println("Name ke :", key, " -->", value)
	}

	//for untuk map
	person := make(map[string]string)
	person["Name"] = "Teguh"
	person["Umur"] = "22"

	for key, value := range person{
		fmt.Println("index -->", key, "-->", value)
	}
}