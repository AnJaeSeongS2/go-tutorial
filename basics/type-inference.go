package basics

import (
	"fmt"
)

func TourTypeInference() {

	// 숫자 상수에 대해 이처럼 처리됨.
	v := 42     // 오른쪽 값으로 타입유추가 됨. // int
	v2 := 42.0  // 오른쪽 값으로 타입유추가 됨. // float64
	v3 := 42.0i // 오른쪽 값으로 타입유추가 됨. // complex128

	fmt.Printf("v(%v) is of Type %T\n", v, v)
	fmt.Printf("v2(%v) is of Type %T\n", v2, v2)
	fmt.Printf("v3(%v) is of Type %T\n", v3, v3)
}
