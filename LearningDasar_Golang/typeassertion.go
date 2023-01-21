/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/
/*

*/
package main

import "fmt"

func random() interface{}{
	return "1"
}
func main(){
	var hasilResult interface{} = random()
	var resultString int = hasilResult.(int)
	fmt.Println(resultString)
}
