package flowcontrol

import (
	"fmt"
)

func TourFor() {
	fmt.Println(sumAs45UsingFor())
	fmt.Println(sumAsGreaterOrEqual1000UsingFor())
	fmt.Println(sumAsGraterOrEquals1000UsingForWithoutSemiColon())
}

func sumAs45UsingFor() any {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func sumAsGreaterOrEqual1000UsingFor() any {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	return sum
}

func sumAsGraterOrEquals1000UsingForWithoutSemiColon() any {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	return sum
}

func foreverUsingFor() {
	for {
	}
}

func CustomSqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}
