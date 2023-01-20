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

type FilterUser func(int, string)bool
func userCheck(idUser int, nameUser1 string, fu FilterUser){
	if fu(idUser) && fu(nameUser1){
		fmt.Println("User dan ID Tersedia: ", idUser, "<->", nameUser1)
	}else{
		fmt.Println("User dan ID tidak ditemukan!")
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
	for i := 0; i < 20; i++{
		fmt.Print("-")
	}

	fu := func(idUser int, nameUser1 string)bool{
		return idUser == 0001 && nameUser1 == "Teguh"
	}
	userCheck(0001,"Teguh",fu)
}