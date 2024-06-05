package main

import "fmt"

/*SWITCH*/

// func menu(fullname string){
// 	switch fullname {
// 	case "Jonathan Farrel Emanuel":
// 		fmt.Println("halo bro waddap")

// 	case "Budi aja":
// 		fmt.Println("halo bro ga kenal")

// 	case "Sudarmono":
// 		fmt.Println("halo pak")

// 	default:
// 		fmt.Println("ga kenal gw skip")

// 	}
// }

/*SWITCH SHORT*/
/*Sama dengan If, Switch juga mendukung short statement sebelum variable yang akan di cek kondisinya*/

func menu(fullname string){

	switch length:= len(fullname); length>5 {
	case true:
		fmt.Println("nama terlalu panjang")
	
	case false:
		fmt.Println("Nama sudah benar")
	}
}


/*Switch without condition*/

// Kondisi di switch expression tidak wajib
// Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case nya
func menu2(fullname string){
	length:=len(fullname)
	switch {
	case length>5:		fmt.Println("Sudah bagus")

	case length>14: 		fmt.Println("nama terlalu panjang")

	case length>27:		fmt.Println("kepanjangan")

	case length>39:		fmt.Println("tau ah skip")
	default:
		fmt.Println("nama sudah benar")

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
// 	menu(fullname)
// 	menu2(fullname)

// }
