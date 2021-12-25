package moretypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	assert.Equal(t, map[string]int{
		"a":   1,
		"b":   1,
		"cc":  2,
		"ddd": 1,
	}, WordCount("a b cc cc ddd"))
}
