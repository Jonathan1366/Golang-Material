package main
import "fmt"

type Filter func(string) string

func sayHi(name string, filter Filter){
	filtered:=filter(name)
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