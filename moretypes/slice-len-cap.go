package moretypes

import (
	"fmt"
)

func TourSliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printMetaSlice(s)

	// slice the slice to give it zero length
	s = s[:0]
	printMetaSlice(s)

	// Extend it's length
	s = s[:4]
	printMetaSlice(s)

	// Extend it's length
	s = s[:6]
	printMetaSlice(s)
	s[4] = 44444

	// Drop its first two values
	s = s[2:]
	printMetaSlice(s)

	s = s[:4]
	printMetaSlice(s)
}

func printMetaSlice(s []int) {
	fmt.Println(getMetaSlice(s))
}

func getMetaSlice(s []int) string {
	return fmt.Sprintf("len=%d cap=%d %v", len(s), cap(s), s)
}
