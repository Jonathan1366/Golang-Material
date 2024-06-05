// package main

// import "fmt"

// func nilai(uas int, uts int)int{
// 	return uas/uts
// }

// func main(){
// 	akhir:=nilai(90,3)
// 	fmt.Println("jadi nilai anda adalah",akhir)
// }


// returning multiple values

// package main
// import "fmt"

// func nama()(string, string){
// 	return "Jonathan", "farrel"
// }

// func main(){
	
// 	// nama1, nama2:=nama()
// 	// fmt.Println(nama1, nama2)

// 	//ignore return value
// 	nama1, _ := nama()
// 	fmt.Println(nama1)

// }


// return values named
// package main
// import "fmt"

// func nama()(depan, belakang,)


package main
import "fmt"

func getCompleteName()(firstName, secondName, lastName string){
	firstName="Jonathan"
	secondName="Farrel"
	lastName="Emanuel"
	fmt.Print()

	return firstName, secondName, lastName
}

// func main(){
// 	a,b,c:=getCompleteName()
// 	fmt.Println(a,b,c)
// }