package main

import "fmt"

// function as parameters
func sayHi(name string, Filter func(string) string){
	filtered:=Filter(name)
	fmt.Println("Halo", filtered)
	
}

func filterWords(name string) string{
	if name=="Anjing"{
		return "..."
	} else{
		return name
	}
}

func main()  {
	var nama string
	fmt.Println("What is your name ?")
	fmt.Scanln(&nama)
	sayHi(nama,filterWords)
	sayHi("Adi", filterWords)


}