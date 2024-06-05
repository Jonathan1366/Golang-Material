package main

import (
	"fmt"
	"hello1/jonathan"
)

func main() {
	result:= jonathan.SayHello("Jonathan")
	fmt.Println(result)

	fmt.Println(jonathan.Os)
	fmt.Println(jonathan.SayGoodbye(jonathan.Os))

}