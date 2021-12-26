package methods

import "fmt"

type I3 interface {
	M3()
}

type T3 struct {
	S3 string
}

func (t *T3) M3() {
	if t == nil {
		fmt.Println("<nil>(custom output)")
		return
	}
	fmt.Println(t.S3)
}

func describe3(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func TourInterfaceValuesWithNil() {
	var i I3
	var t *T3

	fmt.Println(i)
	i = t
	describe3(i) // (<nil>, *methods.T3)
	i.M3()

	i = &T3{"hello"}
	describe3(i)
	i.M3()

	var iNo I3
	describe3(iNo) // (<nil>, <nil>) so, can't call M3 on nextline code.
	// iNo.M3() //runtime exception
}
