// Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
// Struct biasanya representasi data dalam program aplikasi yang kita buat
// Data di struct disimpan dalam field
// Sederhananya struct adalah kumpulan dari field

package main

import "fmt"

type Customer struct{
	name, address string
	age int
}

// func method
func (customer Customer) sayHi(name string){
 fmt.Println("Hello",name,"my name is",customer.name)
}

func main(){
	var Jonathan Customer
	Jonathan.name="Jonathan Farrel Emanuel"
	Jonathan.address="Jakarta"
	Jonathan.age=20

	fmt.Println(Jonathan)
	fmt.Println(Jonathan.name)
	fmt.Println(Jonathan.address)
	fmt.Println(Jonathan.age)

// STRUCT KOMPOSIT

// Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct

nathan:= Customer{
	name:"nathan",
	address: "4pti1",
	age: 20,
}

fmt.Println(nathan)

alvin:= Customer{"alvin","jakarta",30}


fmt.Println(alvin)

alvin.sayHi("tono")


}




