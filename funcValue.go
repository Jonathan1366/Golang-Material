package main

import "fmt"

func hitungAVG(number ... int) int{
	total:=0
	for _, i := range number{
		total+=i
	}
	if len(number)==0{
		return 0 // avoid division with zero
	}
	return total/len(number)
}

func main(){

	var numbers[] int
	var number int
	
	fmt.Println("Masukkan nilai anda:")

	for{
		_, err:=
		fmt.Scanf("%d", &number)
		if err != nil {
			break // out of loop is type zero
		}
		numbers= append(numbers, number)
	}
	// function as value
	final:= hitungAVG(numbers...)
	
	fmt.Println("Jadi rata2nya adalah:",final)
	
}

