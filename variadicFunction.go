//Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
// Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
// Apa bedanya dengan parameter biasa dengan tipe data Array?
// Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
// JIka parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma

package main

import "fmt"
func average(nilai ...int)int{
	total:=0
	for _, number := range  nilai {
		total+=number
	}
	return total/len(nilai)
}
func main() {
	fmt.Println(average(100,90,80))

	// slice params
	terbaru:=[]int{100,99,98}
	fmt.Println(average(terbaru...))
}

