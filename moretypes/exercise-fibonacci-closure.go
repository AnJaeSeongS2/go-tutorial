package moretypes

import (
	"fmt"
)

func fibonacci() func() int {
	// 사실상 반환형 func() int 에서 cur, next에 대한 변경이 cur, next에 계속 남게되므로, func()은 &cur과 &next를 param으로 넘겨받는 것과 동일한 func이다.
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
