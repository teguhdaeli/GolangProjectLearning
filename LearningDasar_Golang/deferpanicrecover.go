/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

/*
Defer --> function yang bisa kita jadwalkan untuk di eksekusi setelah sebuah function diselesaikan
Panic --> function untuk menghentikan program saat berjalan , namun defer ttp jalan
Recover --> function digunakan untuk menangkap data panic, sedangkan aplikasi tidak akan berhenti

*/
package main

import "fmt"

//membuat defer
func logging(){
	fmt.Println("Selesai memanggil function")
}

func runAplication(userNamed string){
	defer logging() //<<<-- memanggil defer function
	if userNamed == "Teguh"{
		fmt.Println("Selamat datang, ",userNamed )
	}else{
		fmt.Println("Username kamu salah" )
	}
}

//membuat panic and recover
func endApp(){
	message := recover()
	if message != nil{
		fmt.Println("Pesan Error : ", message)
	}
	fmt.Println("Aplikasi Selesai!")
}

func rundApp(error bool){
	defer endApp()
	if error{
		panic("ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
	
}
func main(){
	runAplication("Teguh") //<< pemanggilan defer
	rundApp(false) //<< pemanggilan panic
	fmt.Println("Teguh")
}