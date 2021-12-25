package moretypes

// golang에서 function은 클로저일 수  있음. 클로저는 함수의 외부로부터 오는 변수를 참조하는 함수 값인데, 함수는 참조된 변수에 접근할 수 있으나, 이는 bound됨. java와 다르게 go에서는 final취급은 않음.

import "fmt"

func adder() func(int) int {
	sum := 0
	// 이 sum은 return되는 fun(x int) int 에게 closure에 취급되는 변수로서
	return func(x int) int {
		sum += x
		return sum
	}
}

func TourFunctionClosures() {
	pos, neg := adder(), adder()
	var posReturn = 0
	for i := 0; i < 10; i++ {
		posReturn += 1
		// call by value
		posReturn = pos(i)
		fmt.Println(&posReturn)
		fmt.Println(posReturn)
		fmt.Println(posReturn, neg(-2*i))
	}
}
