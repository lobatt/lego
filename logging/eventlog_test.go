package logging

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEventLogEntry(t *testing.T) {
	entry := newEventLogEntry(nil)
	assert.Equal(t, len(entry.Id), 36, "id should be an UUID V4 string")
	assert.Nil(t, entry.Event, "nil input should generate nil entry.event")
}
