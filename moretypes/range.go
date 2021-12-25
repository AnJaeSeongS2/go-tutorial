package moretypes

import (
	"fmt"
)

var powFromTwo = []int{1, 2, 4, 8, 16, 32, 64, 128}

func TourRange() {
	for i, v := range powFromTwo {
		fmt.Println(getCurrentPowFromTwoIndexAndValue(i, v))
	}
}

func getCurrentPowFromTwoIndexAndValue(i int, v int) string {
	return fmt.Sprintf("2^%d = %d", i, v)
}
