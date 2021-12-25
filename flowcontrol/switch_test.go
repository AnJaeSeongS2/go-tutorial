package flowcontrol

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetGoodHourMessage(t *testing.T) {
	assert.Equal(t, "Good morning!", getGoodHourMessage(time.Date(2021, 12, 25, 9, 30, 20, 0, time.UTC)))
}
