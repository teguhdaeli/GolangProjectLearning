/*
	Author Developer	: Teguh Christianto Daeli
	Location 			: Medan, Sumatera Utara
	WhatsApp			: 082162798664
*/

//Slice tipedata, sebuah potongan dari data array serta mirip dengan array, yang membedakan keduanya adalah ukuran slice bisa berubah 
//slice dan array selalu terkoneksi

package main
import "fmt"

func main(){
	//tipe data slice
	/*
		pointer penunjuk data pertama pada array slice
		length, panjang dari isi slice
		capacity, kapasitas dari slice, sehingga length tidak boleh lebih dari capacity

	*/
	var namaBulan = [...]string{"Januari","Februari","Maret","April","Mei","Juni","Juli","Agustus","September","Oktober","November","Desember"}

	var slice1 = namaBulan[2:9]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //<-- cek panjang slice
	fmt.Println(cap(slice1)) //<-- cek kapasitas slice

	namaBulan[5] = "Update" //<-- ubah isi array
	fmt.Println(slice1)
	//
	newSlice1 := make([]string, 2, 5) //<-- membuat slice dengan cara lain dan mudah
	newSlice1[0] = "Teguh"
	newSlice1[1] = "daeli"

	fmt.Println(newSlice1)
	
	//copy slice
	copySlice1 := make([]string, len(newSlice1), cap(newSlice1))
	copy(copySlice1,newSlice1)
	fmt.Println(copySlice1)

	//membedakan array dan slice
	array2 := [5]int{1,2,3,4,5}
	slice2 := []int{1,2,3,4,5}
	fmt.Println(array2)
	fmt.Println(slice2)



}