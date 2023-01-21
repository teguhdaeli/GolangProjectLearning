/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/
/*
Error interface digunakan untuk sebagai kontrak untuk membuat error  secara otomatis dan tidak manual
*/
package main

import (
	"fmt"
	"errors"
)

func pembagian(nilai int, pembagi float)(int, error){
	if pembagi == 0{
		return 0, errors.New("Pembagi tidak boleh nol")
	}else{
		hasil := nilai/pembagi
		return hasil,nil
	}
}

func main(){

	hasil, err := pembagian(20,3)
	if err == nil{
		fmt.Println("Hasil : ", hasil)
	}else{
		fmt.Println("Err: ", err.Error)
	}
}
