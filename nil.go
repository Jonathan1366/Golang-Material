// Go-Lang bukanlah bahasa pemrograman yang berorientasi objek
// Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
// Contoh di Java ada java.lang.Object
// Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
// Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasi nya
// Interface kosong, juga memiliki type alias bernama any

// Ada banyak contoh penggunaan interface kosong di Go-Lang, seperti :
// fmt.Println(a ...interface{})
// panic(v interface{})
// recover() interface{}
// dan lain-lain

package main

import "fmt"

func oldMap(name string) any{
	if name=="" {
		return nil
	} else {
		return name
	}
}

func newMap(name string) map[string] string{
	if name==""{
		return nil
	}else {
		return map[string]string{
			"Name":name,
		}
	}
}


func main() {

	oldmap:=oldMap("")
	fmt.Println(oldmap)
	if oldmap==nil{
		fmt.Println("data map masih kosong")
	}else {
		fmt.Println(oldmap)
	}

	
	newmap:=newMap("Jonathan2")
	fmt.Println(newmap)

	
}