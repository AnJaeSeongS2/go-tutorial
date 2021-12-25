package moretypes

import (
	"fmt"
)

func TourSliceNil() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
