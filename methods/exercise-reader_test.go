package methods

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyReader_Read(t *testing.T) {
	buffer := make([]byte, 1024)
	NewMyReader("test12345678").Read(buffer)
	assert.Equal(t, "test12345678", string(bytes.Split(buffer, []byte{0})[0]))
}
