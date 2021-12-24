package flowcontrol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTourFor(t *testing.T) {
	assert.Equal(t, 45, sumAs45UsingFor())
}
