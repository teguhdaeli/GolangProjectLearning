/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//fucntion parameter

package main

import "fmt"

func hitungHitungPlus(n1 int, n2 int, hasil int){
	hasil = n1 + n2
	// bisa juga hasil dibuat terpisah dari function, misalnya:
	/*
		hasil := n1 + n2
	*/
	fmt.Println("Penjumlahan : ", n1 , " + ", n2, " = ", hasil)

}

func main(){
	//
	hitungHitungPlus(4, 5, 0)
	// jika hasil dipisah dengan function, maka cetak seperti ini:
	//hitungHitungPlus(4, 5)
	//
	//bisa juga menambahkan atau cetak berulang
	n1 := 7 
	n2 :=5 
	hasil := 0
	hitungHitungPlus(n1, n2, hasil)
}