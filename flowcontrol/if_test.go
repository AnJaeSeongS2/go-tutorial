package flowcontrol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTourIf(t *testing.T) {
	assert.Equal(t, "4", sqrt(16))
	assert.Equal(t, "2i", sqrt(-4))
}

func TestTourIfWithAShortStatement(t *testing.T) {
	assert.Equal(t, float64(9), pow(3, 2, 10))
	assert.Equal(t, float64(20), pow(3, 3, 20))
	assert.Equal(t, float64(27), pow(3, 3, 30))
}
