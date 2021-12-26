package methods

import "fmt"

func TourTypeAssertions() {
	var i any = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string) // return underlyingValue, true
	fmt.Println(s, ok)

	f, ok := i.(float64) // return zeroValue, false
	fmt.Println(f, ok)

	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	f = i.(float64) // panic
	fmt.Println(f)
}
