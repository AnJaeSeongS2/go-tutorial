package methods

import "fmt"

func checkType(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TourTypeSwitches() {
	checkType(21)
	checkType("hello")
	checkType(true)
}
