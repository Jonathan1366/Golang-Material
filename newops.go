package main

import "fmt"

type country struct{
	city, province, country string
}

func main() {

	country1:= &country{} // bisa juga sprti berikut new(country)
	country2:=country1

	country2.country="Amerika"

	fmt.Println(country1)
	fmt.Println(country2)

	
}