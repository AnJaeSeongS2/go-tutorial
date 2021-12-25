package moretypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMetaSlice(t *testing.T) {
	s := []int{2, 3, 5, 7, 11, 13}
	assert.Equal(t, "len=6 cap=6 [2 3 5 7 11 13]", getMetaSlice(s))
}
