/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

/*
function value ,<<-- yaitu bisa kita simpan didalam value
-->> value bisa kita buat atau masukkkan ke variabel.

*/

package main
import "fmt"

func getFamily(nameFamily string)string{
 return "Ini Keluarga : " + nameFamily
}

func main(){
 fmt.Println(getFamily("Teguh")) //<<-- cara pertama

 //cara kedua
 resultNameFamily := getFamily("TeguhDaeli")
 fmt.Println(resultNameFamily)

 //cara ketiga
 resultGetName := getFamily
 result := resultGetName("Robin")
 fmt.Println(result)
}