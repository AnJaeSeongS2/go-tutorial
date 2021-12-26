package methods

import "fmt"

func TourEmptyInterface() {
	var i interface{} // go1.18beta1 added builtin any
	describe4(i)

	i = 42
	describe4(i)

	i = "hello"
	describe4(i)

	var i2 any
	describe5(i2)

	i2 = 42
	describe5(i2)

	i2 = "hello"
	describe5(i2)
}

func describe4(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe5(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}
