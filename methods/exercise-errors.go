package methods

import (
	"fmt"
	"math"
)

type ErrUnderZeroFloat struct {
	Value float64
}

func (err ErrUnderZeroFloat) Error() string {
	return fmt.Sprintf("%v value cannot be under zero.", err.Value)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrUnderZeroFloat{x}
	}
	return math.Sqrt(x), nil
}

func TourExerciseErrors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
