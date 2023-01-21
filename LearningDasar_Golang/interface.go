/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/
/*
Interface -->> yaitu tipe data abstract dia tidak memiliki implementasi langsung
--> memiliki defenisi method
*/
package main

import "fmt"
type hasName interface{
	Getname() string
	GetUmur() int
}

func sayHasname(hn hasName){
	fmt.Println("Hallo,", hn.Getname(), "Umur: ", hn.GetUmur())
}

type Person struct{
	Name string
	Umur int
}

func (person Person) Getname() string{
	return person.Name
}
func (person Person) GetUmur() int{
	return person.Umur
}
//////////////////////////////////////////
//////////////////////////////////////////**
/**
type Animal struct{
	NameAnimal string
}

func (animal Animal) Getname() string{
	return animal.NameAnimal
}
*/
func main(){
	//cara implementasi
	var p Person
	p.Name = "Teguh"
	p.Umur = 22
	sayHasname(p)

	/**
	//animal
	a := Animal{
		NameAnimal: "Doggie",
	}
	sayHasname(a)
	*/

}
