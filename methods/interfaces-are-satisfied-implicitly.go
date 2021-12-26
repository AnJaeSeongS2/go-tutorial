package methods

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func TourInterfacesAreSatisfiedImplicitly() {
	var i I = T{"hello"}
	i.M()
}

//

type I2 interface {
	M2()
}

func (t *T) M2() {
	fmt.Println(t.S)
}

type F float64

func (f F) M2() {
	fmt.Println(f)
}

func describe(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func TourInterfaceValues() {
	var i I2

	i = &T{"hello"}

	// (&{hello}, *methods.T)
	describe(i)
	i.M2()

	i = F(math.Pi)
	// (3.141592653589793, methods.F)
	describe(i)
	i.M2()
}
