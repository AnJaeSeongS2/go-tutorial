package moretypes

import "fmt"

func TourMakingSlices() {
	a := make([]int, 5)
	fmt.Println(getMetaSlice2("a", a))

	b := make([]int, 0, 5)
	fmt.Println(getMetaSlice2("b", b))

	c := b[:2]
	fmt.Println(getMetaSlice2("c", c))

	d := b[2:5]
	fmt.Println(getMetaSlice2("d", d))
}

func getMetaSlice2(s string, x []int) string {
	return fmt.Sprintf("%s len=%d cap=%d %v", s, len(x), cap(x), x)
}
