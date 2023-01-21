/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

/*
Closure adalah kemampuan sebuah function berinteraksi dengan data data disekitarnya dalam 
scope yang sama
*/

package main
import "fmt"

func main(){
	i := 0
	//
	increment := func(){ //<<-- cara membuat closure
		fmt.Println("Teguh")
		i++
	}
	//
	increment()
	fmt.Println(i)
}