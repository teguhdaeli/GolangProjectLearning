/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//break & continue --> digunakan untuk menghentikan seluruh perulangan
//continue menghentikan saaat itu sedang berjalan

package main

import "fmt"

func main(){
	//for loop menggunakan break
	for i := 0; i < 10; i++{
		if i == 5{
			break
		}
		fmt.Println("Index Break : ", i)
	}

	//for loop continue
	for i := 0; i < 10; i++{
		if i%2 == 1{ // <-- 1 jika menampilkan genap, 0 jika menampilkan ganjil
			continue
		}
		fmt.Println("Index Continue : ", i)
	}
}