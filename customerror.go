// Untuk membuat error, kita tidak perlu manual.
// Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (Package akan kita bahas secara detail di materi tersendiri)
package main

import (
	"fmt"
)

type validationerror struct{
	Messagge string
}
func (v*validationerror)Error() string {
	return v.Messagge
}

type notfounderror struct{
	Message string
}

func (n *notfounderror)  Error() string{
	return n.Message
}

func SaveData(id string, data any) error{
	if id==""{
		return &validationerror{Messagge: "validation error"}
	}
	if id!="jonathan" {
		return &notfounderror{Message: "Data not found"}
	}
	return nil
}

// cek jenis error
func main() {
	err:=SaveData("jonathan", nil) // ("", nil)
	if err!=nil{
	// 	// terjadi error
	// 	if validationerror, ok:=err.(*validationerror); ok{
	// 		fmt.Println("validation error:", validationerror.Messagge)
	// 	} else if notfounderror, ok:= err.(*notfounderror); ok{
	// 		fmt.Println("not found error:", notfounderror.Message)
	// 	}else {
	// 		fmt.Println("uknown error:", err.Error())
	// 	}
	// }else{
	// 	fmt.Println("sukses")
	// }

	switch finalError:=err.(type) {
	case *validationerror:
		fmt.Println("validation error:", finalError.Error())
	case *notfounderror:
		fmt.Println("not found error:",finalError.Error())
  default:
		fmt.Println("Unknown error:", finalError.Error())
	}
}
}