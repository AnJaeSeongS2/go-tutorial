package basics

import (
	"fmt"
)

const PI = 3.14

func TourConstants() {
	const World = "world"
	// can't use :=
	fmt.Println("hello", World)
	fmt.Println("happy", PI, "day")

	const True = true
	fmt.Println("Go Rules?", True)
}
