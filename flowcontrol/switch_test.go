package flowcontrol

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetGoodHourMessage(t *testing.T) {
	assert.Equal(t, "Good morning!", getGoodHourMessage(time.Date(2021, 12, 25, 9, 30, 20, 0, time.UTC)))
}
