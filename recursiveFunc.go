// Recursive function adalah function yang memanggil function dirinya sendiri
// Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
// Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Factorial

package main

import "fmt"

// func FactorialLoop(nilai int)int {
// 	result:=1
// 	for i:=nilai;i>0; i--{
// 		result*=i
// 	}
// 	return result

// }

func recursiveFactorialLoop(nilai int) int{
	
	if nilai==1 {
		return 1
	}else {
		return nilai*recursiveFactorialLoop(nilai-1)
	}
}

func main(){
	fmt.Println(recursiveFactorialLoop(10))
}