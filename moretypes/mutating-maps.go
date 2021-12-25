package moretypes

import "fmt"

func TourMutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The map:", m)
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	v, ok := m["Answer"]
	// ok is declaration 여부.
	fmt.Println("The map:", m)
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "Answer")
	fmt.Println("The map:", m)
	fmt.Println("The value:", m["Answer"])

	// ok is declaration 여부.
	v, ok = m["Answer"]
	fmt.Println("The map:", m)
	fmt.Println("The value:", v, "Present?", ok)
}
