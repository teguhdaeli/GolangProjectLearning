/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/
/*
Nil atau null >> hanya bisa digunakan dibebrapa tipe data, seperti interface, function, map
slice, pointer dan channel
*/
package main

import "fmt"

func newMap(name string)map[string]string{
	if name == ""{
		return nil
	}else{
		return map[string]string{
			"Name": name,
		}
	}
}

func main(){
	var person map[string]string = newMap("a")
	if person == nil{
		fmt.Println("Data Kosong")
	}else{
		fmt.Println(person)
	}
}
