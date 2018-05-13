package logging

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEventLogEntry(t *testing.T) {
	entry := newEventLogEntry(nil)
	assert.Equal(t, len(entry.id), 36, "id should be an UUID V4 string")
	assert.Nil(t, entry.event, "nil input should generate nil entry.event")
}
