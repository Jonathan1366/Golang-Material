// Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
// Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
package main

import "fmt"

func random() any{
	// return "OK"
	// return 100
	return true

}

func main()  {
	
	// var result any = random()
	// finalResult, ok := result.(string)
	// if !ok{
	// 	fmt.Println("Type assertion failed")
	// }else {
	// 	fmt.Println(finalResult)
	// }

	// Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
	// Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
	// Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
	
// pakai switch
result := random()
switch value:=result.(type) {
case string:
	fmt.Println("String",value)
case int:
	fmt.Println("Integer",value)
default:
	fmt.Println("Tidak diketahui")
}
	
}



