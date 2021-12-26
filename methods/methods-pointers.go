package methods

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs3() float64 {
	fmt.Println(&v)
	fmt.Println(v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs3(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex3) Scale(f float64) {
	// struct 에 대해서는 pointer를 뱉으면 &value로 표기된다.
	// struct 에 대해서는 pointer에서 바로 dot을 통해 field를 사용가능하다.
	fmt.Println(v)
	fmt.Println(*v)
	v.X = v.X * f
	v.Y = v.Y * f
}

// if Vertex3 without *, copy by value
func Scale(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func PointerCheck(v *Vertex3, v2 Vertex3) {
	// v2 is new value
	fmt.Println(v == &v2)
	fmt.Println(*v == v2)
}

// value receiver case
func (v Vertex3) PointerCheck2(v2 *Vertex3) {
	fmt.Println(&v == v2)
	v.Scale(2)
}

// pointer receiver case
// 1. receiver 값을 수정하기 원하거나
// 2. receiver가 복사되는 것이 무거울 경우, 이를 채택가능.
// 한 type/ struct에 대한 receiver운용에는 value/pointer 둘중 하나를 택해 일관성을 유지할 것...
func (v *Vertex3) PointerCheck3(v2 *Vertex3) {
	fmt.Println(v == v2)
	v.Scale(2)
}

func TourMethodsPointers() {
	v := Vertex3{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs3())

	v = Vertex3{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs3(v))

	PointerCheck(&v, v)
	fmt.Println(v)
	fmt.Println("PointerCheck2 start")
	v.PointerCheck2(&v)
	fmt.Println(v)
	fmt.Println("PointerCheck3 start")
	v.PointerCheck3(&v)
	fmt.Println(v)

	v = Vertex3{3, 4}
	v.Scale(2)
	Scale(&v, 10)

	p := &Vertex3{4, 3}
	p.Scale(3)
	Scale(p, 8)

	fmt.Println(v, p)
}
