package logging

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"os"
	"time"
)

var eventLog = log.New(os.Stdout, "", 0)

type EventLogEntry struct {
	id        string
	timestamp int64
	event     interface{}
}

func newEventLogEntry(e interface{}) *EventLogEntry {
	entry := &EventLogEntry{id: uuid.NewV4().String(), timestamp: time.Now().UnixNano(), event: e}
	return entry
}

// SetOutput set the output io.Writer for global event logger
func SetOutput(w io.Writer) {
	eventLog.SetOutput(w)
}

// EventLog logs an event defined by caller in json format
func LogEvent(event interface{}) {
	entry := newEventLogEntry(event)
	logJson, e := json.Marshal(entry)
	if e != nil {
		eventLog.Printf("{error: %s}\n", e.Error())
		return
	}

	eventLog.Printf("%s\n", logJson)
}
