package methods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertex_Abs(t *testing.T) {
	assert.Equal(t, float64(13), Vertex{12, 5}.Abs())
}
