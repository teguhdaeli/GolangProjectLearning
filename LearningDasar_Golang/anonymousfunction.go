/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

/*
Function Anonymous adalah function yang tidak memiliki nama atau identitas.


*/

package main
import "fmt"

type Blacklist func(string)bool
func registerUser(nameUser string, bl Blacklist){
	if bl(nameUser){
		fmt.Println("Kamu diblok, ", nameUser)
	}else{
		fmt.Println("Selamat datang, ", nameUser)
	}
}

func main(){
	//cara pertama buat panggil parameternya
	bl := func(nameUser string)bool{
		return nameUser == "Admin"
	}
	registerUser("Admin", bl)
	registerUser("Teguh", bl)

	//cara kedua buat panggil
	registerUser("admin", func(nameUser string)bool{
		return nameUser == "admin"
	})

	/*
		Jika gunakan cara pertama, kita hanya panggil sekali dan bisa cetak berkali kali
		Jika gunakan cara kedua, kita panggil berulang kali function dan outpuntnya
	*/
}