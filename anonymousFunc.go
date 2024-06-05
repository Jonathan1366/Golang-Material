// Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
// Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
// Hal tersebut dinamakan anonymous function, atau function tanpa nama

package main

import "fmt"

type Blacklist func(string) bool

func registUsername(name string, blokir Blacklist)  {
	if blokir(name) {
		fmt.Println("You're blocked", name)
	} else {
		fmt.Println("welcome",name)
	}
	
}

func main()  {
	// anonymus func
	terfilter:=func(name string) bool{
		return name== "anjing"
	}
	registUsername("jonathan", terfilter)
	registUsername("anjing", terfilter)

	
}