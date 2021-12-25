package moretypes

import (
	"fmt"
)

func TourPointers() {
	i := 42
	p := &i
	fmt.Printf("%v : %v\n", *p, p)
	*p = 21
	fmt.Printf("%v : %v\n", *p, p)

	i = 42
	j := 2701

	p = &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
