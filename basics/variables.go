package basics

// var 문은 package / function 단에서 존재 가능.
// 선언만으로 기본값이 존재한다. primitive 값이기 때문인듯..?

import (
	"fmt"
)

var c, python, java bool

func TourVariables() {
	var i int
	fmt.Println(i, c, python, java)
}
