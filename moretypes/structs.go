package moretypes

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	staticV1 = Vertex{1, 2}
	staticV2 = Vertex{Y: 1}
	staticV3 = Vertex{}
	staticP  = &Vertex{4, 3} // &{4 3} 으로 찍힘
	staticP2 = &staticV2     // &{0 1} 으로 찍힘
)

func TourStruct() {
	vertex := Vertex{1, 2}
	fmt.Println(vertex)
	vertex.X = 4
	fmt.Println(vertex)

	p := &vertex
	(*p).X = 1e9
	fmt.Println(vertex)
	p.X = 1e6 // struct에 대한 pointer의 경우, struct's field에 대한 접근 을 위한 struct pointer에 대한 dot . 표기는 (*p).X임을 의미하게 된다.
	fmt.Println(vertex)

	fmt.Println(staticV1, staticV2, staticV3, staticP, staticP2)
}
