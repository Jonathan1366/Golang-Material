package main

// interface statis
import "fmt"

type Geometry interface{
	Area() float64
}


type Triangle struct{
	Base float64
	Height float64
}

func (t Triangle) Area() float64{
	return 0.5*t.Base*t.Height
}


// ukur
func Measure (g Geometry){
	fmt.Println("Area:",g.Area())
}

func main() {
	t:= Triangle {Base: 3, Height: 4}
	Measure (t)
}
