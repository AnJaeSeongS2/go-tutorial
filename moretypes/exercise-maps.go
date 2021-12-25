package moretypes

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, v := range strings.Fields(s) {
		result[v] += 1
	}
	return result
}

func TourExerciseMaps() {
	wc.Test(WordCount)
}
