// Go-Lang bukanlah bahasa pemrograman yang berorientasi objek
// Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
// Contoh di Java ada java.lang.Object
// Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
// Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasi nya
// Interface kosong, juga memiliki type alias bernama any

package main

import "fmt"

func tambah (v interface{}){
	fmt.Println(v)
}


func kurang () any{
	return 1 
}

func main() {
	tambah("halo")
	tambah(64)
	tambah(true)

	var kosong any= kurang()
	fmt.Println(kosong)


}
