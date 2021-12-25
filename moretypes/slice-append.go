package moretypes

import "fmt"

func TourSliceAppend() {
	var s []int
	fmt.Println(getMetaSlice3(s))

	s = append(s, 0)
	fmt.Println(getMetaSlice3(s))

	s = append(s, 1)
	fmt.Println(getMetaSlice3(s))

	s = append(s, 2, 3, 4)
	fmt.Println(getMetaSlice3(s))
}

func getMetaSlice3(s []int) string {
	return fmt.Sprintf("len=%d cap=%d %v", len(s), cap(s), s)
}
