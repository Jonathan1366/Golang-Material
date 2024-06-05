// Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error
// Untuk membuat error, kita tidak perlu manual.
// Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (Package akan kita bahas secara detail di materi tersendiri)

package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error){
	if pembagi==0 {
		// return 0, fmt.Errorf("Nilai=%d tidak bisa dibagi dengan pembagi=%d", nilai, pembagi)
		return 0, errors.New("Tidak bisa di bagi dengan 0")
	} else {
		return nilai/pembagi, nil
		// return hasil nya, errornya tinggal diberikan nil
	}
}

func main() {
	hasil, err:=pembagian(100,0) // (100,0)
	if err==nil {
		fmt.Println("Hasil", hasil)
	}else{
		fmt.Println("Error", err.Error())
	}
}


