package logging

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"regexp"
	"testing"
)

var UUIDV4Pattern = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"
var TimestampPattern = "[0-9]{19}"

type TestEvent struct {
	URL          string
	Method       string
	Body         []byte
	ResponseCode int
}

var testEvent = &TestEvent{URL: "http://github.com", Method: "POST", Body: []byte("Hello"), ResponseCode: 200}

func ExampleLogEvent() {
	SetOutput(os.Stdout)
	LogEvent(testEvent)
}

func TestNewEventLogEntry(t *testing.T) {
	entry := newEventLogEntry(nil)
	assert.Regexp(t, regexp.MustCompile(UUIDV4Pattern), entry.Id, "id should be an UUID V4 string")
	assert.Nil(t, entry.Event, "nil input should generate nil entry.event")
}

func TestLogEvent(t *testing.T) {
	var buf []byte
	var b = bytes.NewBuffer(buf)
	SetOutput(b)
	LogEvent(testEvent)

	assert.Regexp(t,
		regexp.MustCompile(
			fmt.Sprintf(`{"id":"%s","timestamp":%s,"event":{"URL":"http://github.com","Method":"POST","Body":"SGVsbG8=","ResponseCode":200}}\n`, UUIDV4Pattern, TimestampPattern)), b.String())
}

func TestLogEventMultiple(t *testing.T) {
	var buf []byte
	var b = bytes.NewBuffer(buf)
	SetOutput(b)
	LogEvent(testEvent)
	LogEvent(testEvent)

	line1, e1 := b.ReadString('\n')
	assert.Nil(t, e1)
	assert.Regexp(t,
		regexp.MustCompile(
			fmt.Sprintf(`{"id":"%s","timestamp":%s,"event":{"URL":"http://github.com","Method":"POST","Body":"SGVsbG8=","ResponseCode":200}}`, UUIDV4Pattern, TimestampPattern)), line1)
	line2, e2 := b.ReadString('\n')
	assert.Nil(t, e2)
	assert.Regexp(t,
		regexp.MustCompile(
			fmt.Sprintf(`{"id":"%s","timestamp":%s,"event":{"URL":"http://github.com","Method":"POST","Body":"SGVsbG8=","ResponseCode":200}}`, UUIDV4Pattern, TimestampPattern)), line2)
}
