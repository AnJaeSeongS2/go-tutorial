package basics

// var문 사용시, 초깃값이 존재하면, type은 생략가능.

import (
	"fmt"
)

var i, j int = 1, 2

func TourVariablesWithInitializers() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
