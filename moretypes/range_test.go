package moretypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentPowFromTwoIndexAndValue(t *testing.T) {
	assert.Equal(t, "2^2 = 4", getCurrentPowFromTwoIndexAndValue(2, 4))
}
