package main

import (
	"fmt"
)

func variable() {
	// print integer
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)
	fmt.Print("Tiga koma lima=", 3.5)

	// print boolean
	fmt.Println("Benar=", true)
	fmt.Println("Salah =", false)

	fmt.Println(len("halo"))
	fmt.Println("Jonathan Farrel"[0])
	fmt.Println("Jonathan Farrel"[2])

	// deklarasi variable dan call var
	var name string
	var angka = 12
	fmt.Println(angka)
	var angka1 = true
	fmt.Println(angka1)
	var terbaru = 1.34
	fmt.Printf("%.2f\n", terbaru) // untuk pindah baris baru pakai %.2f\n

	// print name
	name= "Jonathan Farrel Emanuel"
	fmt.Println("jadi nama saya adalah ",name)
	name="nathan"
	fmt.Println("jadi panggilan saya adalah", name)
	

	//keyword var
	// short declaration
	str1 := "hai"
	str2 := "aku"

	fmt.Print(str1,str2)
	fmt.Print()

	// multiple variable
	var (
		nama1 = "jonathan"
		nama2 = "farrel"
		nama3 = "emanuel"
	)
	fmt.Println(nama1, nama2, nama3)

	// const & multiple const
	const a = 1
	const b = 2
	fmt.Println(a, b)

	const (
		hai1 = 2
		hai2 = 10
	)
	fmt.Println(hai1, hai2)
	fmt.Println(hai1)
	fmt.Println(hai2)

	// conversion data type
	var nilai32 int32=32768
	var nilai64 int64=int64(nilai32)

	var nilai16 int16 =int16(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) // number overflow karena melebihi kapasitas

	// conversion data type 2
	// change index chr to str

	name23:="jonathanaja"
	fmt.Println(name23[0])

	name23="jonathanaja"
	var e = name23[0] // var == uint8
	var eStr= string(e)

	fmt.Println(eStr) // show the real chr

	// TYPE DECLARATION
	/*Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti
*/

	type ktp string
	var ktpJo ktp="1234567890"
	fmt.Println(ktpJo)

	// convert str alias
	var contohKTP string ="22222"
	var ktporang ktp= ktp(contohKTP)
	println(ktporang)  

	// MATH OPS

	var abaru=10
	var a1=10
	var a2=10
	var a3=10

	print("hasil dari abaru + a1 - a2 x a3 = ", abaru+a1-a2*a3)
	fmt.Println()

	// augmented assignment

	// a=a+10

	var i = 10
	i += 10
	fmt.Println(i)

	// unary operator
	// shortcut
	var j=1
	j++
	println(j)
	j++
	println(j)

	// operasi perbandingan
	/*Operasi perbandingan adalah operasi untuk membandingkan dua buah data
Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
Jika hasil operasinya adalah benar, maka nilainya adalah true
Jika hasil operasinya adalah salah, maka nilainya adalah false
*/

var nama11="Jonathan"
var nama22="farrel"

var result bool= nama11 == nama22
fmt.Println(result)

nama11="jonathan"
nama22="jonathan"
var result1 bool= nama11 == nama22
var result12 bool= nama11 != nama22

fmt.Println(result1,result12)

// BOOLEAN OPS

// compare nilai uas dan uts, jika 22nya lebih dari 80 maka lulus 

var UAS = 90
var uts= 80

var lulusUAS bool = UAS>80
var lulusUTS bool = uts>80

var lulus bool= lulusUAS&& lulusUTS

// true= lulus false= ga lulus
fmt.Println(lulus)


// ARRAY
/*Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
Daya tampung Array tidak bisa bertambah setelah Array dibuat
*/

var names [3] string // tidak bisa di isi langsung [...]
names [0]="Jonathan"
names [1]="Farrel"
names [2]="Emanuel"

fmt.Println(names[0],names[1],names[2])

var arraynya= [3]int{
	1,2,3,
}

fmt.Print(arraynya)
fmt.Println(arraynya)


// func array len(), array[index] => akses index, array[index]=0 ==> change index values

// harus di isi langsung
var arraynya2= [...]int{
	1,2,3,
}
fmt.Println(len(arraynya2))


// SLICE DATA

/*Tipe data Slice adalah potongan dari data Array
Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
Slide dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array


Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
Pointer adalah penunjuk data pertama di array para slice
Length adalah panjang dari slice, dan
Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity


array[low:high]
Membuat slice dari array dimulai index low sampai index sebelum high
array[low:]
Membuat slide dari array dimulai index low sampai index akhir di array
array[:high]
Membuat slice dari array dimulai index 0 sampai index sebelum high
array[:]
Membuat slice dari array dimulai index 0 sampai index akhir di array
*/

fullname:=[...]string{"Jo","nathan","farr","rel","emanuel","huh"}
slices1:=fullname[4:6]
fmt.Println(slices1)
// fmt.Println(slices1)

slices2:= fullname[:3]
fmt.Println(slices2)
slices3:= fullname[3:]
fmt.Println(slices3)
slices4:= fullname[:]
fmt.Println(slices4)


/*function slice


len(slice)
Untuk mendapatkan panjang 
cap(slice)
Untuk mendapat kapasitas 
append(slice, data)
Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
make([]TypeData, length, capacity)
Membuat slice baru
copy(destination, source)
Menyalin slice dari source ke destination
*/

	days:=[...]string{"monday","tuesday","wednesday","thursday","friday","saturday","sunday"}
	dayslice1:= days[5:]
	dayslice1[0]="Sabtu baru"
	dayslice1[1]="Minggu baru"

	fmt.Println(days)
	
	dayslice2:= append(dayslice1,"Libur panjang")
	fmt.Println(dayslice1)
	fmt.Println(dayslice2)
	fmt.Println(days)

	// make slice == make([]TypeData, length, capacity)

	var sliceBaru = make([]string, 2,5)
	sliceBaru[0]="Jonathan"
	sliceBaru[1]="Jonathan"
	// tidak bisa tambah manual hrus pakai append

	sliceterbaru:= append(sliceBaru, "apakah")
	
	fmt.Println(sliceBaru)
	fmt.Println(cap(sliceBaru))
	fmt.Println(len(sliceBaru))
	fmt.Println(sliceterbaru)


	nilaiTerbaru:= [...] int8{90,91,92,93,94,95,96,97,98}
	fmt.Println(nilaiTerbaru)

	newSLICE:= nilaiTerbaru[4:8]
	fmt.Println(newSLICE)

	addSLICE:=append(newSLICE,100)
	fmt.Println(addSLICE)
	fmt.Println(len(addSLICE))
	fmt.Println(cap(addSLICE))


	nilaiTerlama:= [...] int8{80,81,82,83,84,85,86,87,88}
	fmt.Println(nilaiTerlama)

	//copy slice

	fromSlice:= nilaiTerlama[:]
	toSlice:=nilaiTerbaru[:]
	copy(fromSlice,toSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)


	nilaiAJA:= [...] int8{80,81,82,83,84,85,86,87,88}
	fmt.Println(nilaiAJA)


	fromSlice1:= nilaiAJA[:]
	toSlice1:=make([]int8, len(fromSlice1))
	copy(toSlice1, fromSlice1)
	fmt.Println(fromSlice1)
	fmt.Println(toSlice1)

	// slice  vs array
	iniArr:= [...]int8{1,2,3,4,5,6,7,8}
	iniSlice:= []int8{1,2,3,4,5,6,7,8}
	fmt.Println(iniArr)
	fmt.Println(iniSlice)

	// rata2 pakai slice utk di Golang
	

	// map data type
	/*Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru
*/
	person:= map[string]string{
		"user":"Jonathan Farrel Emanuel",
		"NIM":"32220020",
	}

	fmt.Println(person["user"])
	fmt.Println(person["NIM"])
	fmt.Println(person)


	//function in map
	Copilot:= map[string]string{
		"brain":"Deep Neural Network",
		"method":"Transfer Learning",
		"user":"Jonathan Farrel Emanuel",
	}
	

	// menghitung jumlah data dalam map
	fmt.Println("Jumlah data dalam map:",len(Copilot))

	// mengambil nilai berdasarkan key
	fmt.Println("Nilai untuk key 'brain':",Copilot["brain"])
	

	// mengubah nilai berdasarkan key
	Copilot["brain"]= "KNN"

	// menghitung jumlah data setelah berubah
	fmt.Println("Jumlah data setelah berubah adalah: ",len(Copilot))

	// menghapus data berdasarkan key
	delete(Copilot, "user")

	// menghitung jumlah dtata setelah penghapusan
	fmt.Println("Jumlah data setelah penghapusan: ", len(Copilot))


	/* len(map)
Untuk mendapatkan jumlah data di map
map[key]
Mengambil data di map dengan key
map[key] = value
Mengubah data di map dengan key
make(map[TypeKey]TypeValue)
Membuat map baru
delete(map, key)
Menghapus data di map dengan key
*/


Copilot2 := make(map[string]string)
Copilot2["brain"]="Deep Neural Network"
Copilot2["method"]="LLM"
Copilot2["languange"]="Python"

fmt.Println("Map Copilot adalah:", Copilot2)
delete(Copilot2,"languange")
fmt.Println("Map Copilot adalah:", Copilot2)


	



}
// func main(){
// 	variable()
// }


