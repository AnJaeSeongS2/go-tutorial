package flowcontrol

import (
	"fmt"
)

func TourFor() {
	fmt.Println(sumAs45UsingFor())
}

func sumAs45UsingFor() any {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}
