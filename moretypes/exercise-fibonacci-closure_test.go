package moretypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	fibo := fibonacci()
	for i := 0; i < 3; i++ {
		fibo()
	}

	assert.Equal(t, 2, fibo())
}
