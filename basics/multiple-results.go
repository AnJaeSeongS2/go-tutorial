package basics

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func TourMultipleResults() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
