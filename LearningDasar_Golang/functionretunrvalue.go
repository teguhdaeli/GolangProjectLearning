/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//Function return value --> mengembalikan data apa yg kita buat

package main

import "fmt"

func getName(name string)string{
	if name == ""{
		return "Hallo, bro"
	}else if name == "Teguh"{
		return "Hallo, " + name
	}else{
		return "Kenalan Dong : " + name
	}
	//
}
func main()  {
	fmt.Println(getName("Elsa"))
}