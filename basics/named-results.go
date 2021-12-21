package basics

/*
naked return value
여기 이 예제에서처럼 naked return문은 짧은 함수에서만 사용되어야합니다. 긴 함수에서는 그것이 가독성을 떨어뜨릴 수 있습니다.
*/

import (
	"fmt"
)

func split(src int) (x, y int) {
	x = src * 4 / 9
	y = src - x
	return
}

func TourNamedResults() {
	fmt.Println(split(63))
}
