package moretypes

import (
	"fmt"
)

func TourSlices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int
	s = primes[1:4]
	fmt.Println(s)

	// slices-pointers.go

	names := [4]string{
		"jaeseong",
		"jihye",
		"jaesun",
		"johan",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[1] = "EMPTY"
	fmt.Println(a, b, names)

	primesSliceLiterals := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primesSliceLiterals)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	structTest := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(structTest)

	sliceBoundsOrigin := []int{2, 3, 5, 7, 11, 13}
	sliceBounds := sliceBoundsOrigin
	fmt.Println(sliceBounds)

	sliceBounds = sliceBounds[1:4]
	fmt.Println(sliceBounds)

	sliceBounds = sliceBounds[:]
	fmt.Println(sliceBounds)

	sliceBounds = sliceBounds[:2]
	fmt.Println(sliceBounds)

	sliceBounds = sliceBounds[1:]
	fmt.Println(sliceBounds)
	fmt.Println(sliceBoundsOrigin)
}
