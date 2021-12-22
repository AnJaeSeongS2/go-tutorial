package basics

import (
	"fmt"
)

const (
	Big   = 1 << 100  // 1.2676506002282295e+29
	Small = Big >> 99 // 2
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func TourNumericConstants() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Printf("Small type : %T\n", Small)
	// fmt.Println(needInt(Big))
	// basics/numeric-constants.go:23:22: cannot use Big (untyped int constant 1267650600228229401496703205376) as int value in argument to needInt (overflows)
	fmt.Println(needFloat(Big))
	// fmt.Printf("Big type : %T\n", Big)
}
