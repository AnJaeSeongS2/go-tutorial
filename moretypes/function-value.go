package moretypes

//일급 함수에 대한 이야기이다.

import (
	"fmt"
	"math"
)

func compute3And4WithFunction(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func TourFunctionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute3And4WithFunction(hypot))
	fmt.Println(compute3And4WithFunction(math.Pow))
}
