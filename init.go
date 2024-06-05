package main

import (
	"fmt"
	"hello1/database"
	_ "hello1/internal"
)

func main(){
	fmt.Println(database.Getdatabase())
}