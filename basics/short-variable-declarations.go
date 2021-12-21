package basics

// :=은 function 내에서만  사용 제한.
// function 밖에서는 모든 선언이 키워드 (var, func, etc...)로 시작해야함.

import (
	"fmt"
)

func TourShortVariableDeclarations() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
