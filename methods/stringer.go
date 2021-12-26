package methods

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func TourStringer() {
	a := Person{"jaeseong", 31}
	z := Person{"jihye", 30}
	fmt.Println(a, z)
}
