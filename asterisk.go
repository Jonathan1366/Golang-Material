// Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
// Semua variable yang mengacu ke data yang sama tidak akan berubah
// Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *

package main

import "fmt"

type negara struct{
	jumlah_penduduk, bahasa, skill string
}
func main() {
	panggil1:=negara{"20 juta","English","Coding"}
	panggil2:=&panggil1
	panggil2.skill="Entrepreneur"

	fmt.Println(panggil1)
	fmt.Println(panggil2) 

	// panggil2= &negara{"30 juta","Indonesia","Coding"}
	// fmt.Println(panggil1)
	// fmt.Println(panggil2)  
	*panggil2 = negara{"30 juta","Indonesia","Coding"}
	fmt.Println(panggil1)
	fmt.Println(panggil2)  

	

}