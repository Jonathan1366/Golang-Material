// Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
// Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi

package main

import "fmt"

func add(x int) func() int{
	return func() int {
		x+=1
		return x
	}
}

func main(){
	increment:= add(1)
	fmt.Println(increment())
	fmt.Println(increment()) 
	fmt.Println(increment())
}