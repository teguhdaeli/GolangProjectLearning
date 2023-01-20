/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

/*Return multiple value

-->mengembalikan data lebih dari satu
-->cara agar bisa multiple value, kita harus tulis semua return valuenya

*/

package main

import "fmt"
//cara pertama getFullname

func getFullname() (string, string, int){ //<-- multiple value
	return "Teguh", "Daeli", 22
}

func getPenjumlahan() (int, int, int){ //<-- untuk penjumlahan
	return 2, 9, 0
}

func main(){
	// jika kita tidak membutuhkan salah satu variabelnya, kita bisa gunakan "_" atau underscore
	firstName, LastName, Umur := getFullname()
	fmt.Println(firstName, LastName, Umur)
	//
	//untuk perhitungan penjumlahan
	n1, n2, hasil := getPenjumlahan() //<-- misalnya n1, n2, _ := getPenjumlahan() <-- jika kita tidak membutuhkan salah satu variabelnya,
	hasil = n1 + n2
	fmt.Println(n1, n2,"=", hasil)
	//

}