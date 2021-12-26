package methods

import (
	"fmt"
	"math"
)

// _interface_type_ 은 method's signature 집합
type Abser5 interface {
	Abs5() float64
}

func TourInterfaces() {
	var a Abser5
	f := MyFloat5(-math.Sqrt2)
	v := Vertex5{3, 4}

	a = f
	fmt.Println(a.Abs5())

	a = &v
	fmt.Println(a.Abs5())

	// a = v
	fmt.Println(a.Abs5())
}

type MyFloat5 float64

func (f MyFloat5) Abs5() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex5 struct {
	X, Y float64
}

func (v *Vertex5) Abs5() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
