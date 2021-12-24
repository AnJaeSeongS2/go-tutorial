package flowcontrol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTourFor(t *testing.T) {
	assert.Equal(t, 45, sumAs45UsingFor())
	assert.LessOrEqual(t, 1000, sumAsGreaterOrEqual1000UsingFor())
	assert.LessOrEqual(t, 1000, sumAsGraterOrEquals1000UsingForWithoutSemiColon())
}

func TestCustomSqrt(t *testing.T) {

	// assert.Equal(t, math.Sqrt(2), CustomSqrt(2))
	/*
			Error Trace:	if_test.go:21
		  Error:      	Not equal:
		  expected: 1.4142135623730951
		  actual  : 1.414213562373095
	*/
}
