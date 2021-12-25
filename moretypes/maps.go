package moretypes

import (
	"fmt"
)

type VertexOnMaps struct {
	Lat, Long float64
}

var m map[string]VertexOnMaps

var literalMap = map[string]VertexOnMaps{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func TourMaps() {
	m = make(map[string]VertexOnMaps)
	m["Bell Labs"] = VertexOnMaps{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)

	fmt.Println(literalMap)
}
