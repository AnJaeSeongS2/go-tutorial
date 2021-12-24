package flowcontrol

import (
	"fmt"
)

func TourDefer() {
	// defer는 stack으로 쌓이고, stack에서 꺼내와 실행된다. 따라서, LIFO 이다.
	defer fmt.Println(getWorld())
	defer fmt.Println("before defer")
	defer fmt.Println("more before defer")
	fmt.Println(getHello())
}

func getWorld() string {
	return "world!"
}

func getHello() string {
	return "Hello"
}
