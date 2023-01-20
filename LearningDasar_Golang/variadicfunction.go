/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

/*Variadic Function <-- parameter berada pada terkahir, bisa dijadikan sebagai varags yaitu datanya bisa 
menerima lebih dari satu input atau seperti array.
*/

package main
import "fmt"

func sumAll(number ...int)int{
	total := 0
	for _, value := range number{
		total += value
	}
	return total
}

func main()  {
	//cara 
	total := sumAll(5)
	fmt.Println(total)
	//menggunakan slice dengan variadic
	slice := []int{5,5,6,2}
	total = sumAll(slice...)
	fmt.Println(total)
}