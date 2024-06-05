// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
// Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

// defer
// package main

// import "fmt"

// func status(){
// 	fmt.Println("Online")
// }

// func runApp(){
// 	defer status()
// 	fmt.Println("RunApp")
// }

// func main(){
// 	runApp()
// }

// Panic function adalah function yang bisa kita gunakan untuk menghentikan program
// Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
// Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

// package main

// import "fmt"

// func end(){
// 	fmt.Println("End app")
// }

// func Status(status bool){

// 	defer end()
// 	if status{
// 		panic("error anjay")
// 	}
// }

// func main(){
// 	Status(true)
// }

// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

package main

import "fmt"

func endApp(){
	fmt.Println("End App")
	pesan:= recover()
	fmt.Println("terjadi panic", pesan)
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("Ups error anjay")
	}
}

func main(){
	runApp(true)
	fmt.Println("Jonathan Farrel Emanuel")
}
