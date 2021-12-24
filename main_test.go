package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestify(t *testing.T) {
	assert.Equal(t, 10, 10, "equal")
	assert.NotEqual(t, 11, 10, "not equal")
}
