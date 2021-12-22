package basics

// C와 달리 Go는 다른 type의 요소들 간의 할당에는 명시적인 변환을 필요로합니다.
import (
	"fmt"
	"math"
)

func TourTypeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
