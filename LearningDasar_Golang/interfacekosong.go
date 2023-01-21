/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/
/*
Interface kosong yang tidak memiliki deklarasi method satupun
Contoh menggunakan:
fmt.Prinln(bla bla...interface[])
panic(v.interface[])
revocer()interface[]

dan sebagainya......
*/
package main

import "fmt"

func Upps(i interface{})interface{}{
	if i == 1{
		return 1
	}else if i == 2{
		return 2
	}else{
		return "Tidak ada"
	}
}

func main(){
	fmt.Println(Upps(2))
}
