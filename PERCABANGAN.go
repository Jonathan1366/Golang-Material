package main

import "fmt"

/*If adalah salah satu kata kunci yang digunakan untuk percabangan
Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
Hampir di semua bahasa pemrograman mendukung if expression
*/

/*Blok if akan dieksekusi ketika kondisi if bernilai true
Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
Hal ini bisa dilakukan menggunakan else expression
*/

/*If mendukung short statement sebelum kondisi
Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi
*/

// func percabangan(fullname string, shortname string ){

// 	if fullname== "Jonathan Farrel Emanuel" {
// 		fmt.Println("Halo jonathan")
// 	} else if shortname=="Nathan" {
// 		fmt.Println("halo Nathan")
// 	} else{
// 		fmt.Println("who are you ?")
// 	}
// }

func percabangan(fullname string, shortname string ){
	
	if fullname== "Jonathan Farrel Emanuel"||shortname=="Nathan" {
		fmt.Println("Halo bestie")
	} else{
		fmt.Println("who are you ?")
	}
}


// func main(){

	
// 	fmt.Println("Enter your firstname: ")
// 	var firstname string

// 	fmt.Scanln(&firstname)
// 	fmt.Println("Enter your last name: ")
// 	var secondname, lastname string
// 	fmt.Scanln(&secondname,&lastname)	
// 	var shortname string
// 	fmt.Println("enter your shortname: ",shortname)
// 	fmt.Scanln(&shortname)	

// 	fullname:= firstname +" "+ secondname +" "+ lastname
// 	fmt.Println("Your full name is: ", fullname)
// 	fmt.Println("Your shortname is: ", shortname)
// 	percabangan(fullname, shortname)

// }
