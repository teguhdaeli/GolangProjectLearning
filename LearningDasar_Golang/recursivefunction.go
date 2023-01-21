/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

/*Recursive function --> function yang memanggil dirinya sendiri
Contoh kasus yang bisa diselesaikan --> factorial
*/

package main
import "fmt"

//cara pertama
func factorialLoop(value int)int{
	result := 1
	for i := value; i > 0; i--{
		result *= i
	}
	return result
}

//cara kedua
func factorialRecursive(n1 int)int{
	if n1 == 1{
		return 1
	}else{
		return n1 * factorialRecursive(n1-1)
	}
}


func main(){
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursive(1))
}