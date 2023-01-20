/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//Switch logika

package main
import "fmt"

func main(){
	//cara pertama pakai variabel di awal
	name := "Daeli"

	switch name{
	case "Teguh":
		fmt.Println("Hello, ", name)
	
	case "Dona":
		fmt.Println("Hallo, Dona")
	
	default:
		fmt.Println("Kenalan dong, ", name)
	}

	//cara kedua short statemen tanpa variabel di awal

	switch umur := 29; umur > 31{
	case true:
		fmt.Println("Halo umur, ", umur)
	case false:
		fmt.Println("Umur kamu,", umur, " ya")
	default:
		fmt.Println("Umur kamu muda")
	}

	// switch tanpa kondisi
	nameLong := "Teguh dale"
	numName := len(nameLong)

	switch{
	case numName > 10:
		fmt.Println("Nama Panjang ya")
	case numName > 5:
		fmt.Println("Nama Lumayan panjang ya")
	default:
		fmt.Println("Hmm panjang ya!!")
	}
}