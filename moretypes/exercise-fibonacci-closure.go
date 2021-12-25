package moretypes

import (
	"fmt"
)

func fibonacci() func() int {
	cur, next := 0, 1
	return func() int {
		prev := cur
		cur = next
		next = prev + cur
		return prev
	}
}

func TourExerciseFibonacciClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
