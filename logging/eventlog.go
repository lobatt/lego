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
	Id        string      `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Event     interface{} `json:"event"`
}

// LoggableEvent is a interface with only one function `ToLogRecord`
// which let user customize the information they want to put in event logging.
type LoggableEvent interface {
	ToLogRecord() string
}

func newEventLogEntry(e interface{}) *EventLogEntry {
	entry := &EventLogEntry{Id: uuid.NewV4().String(), Timestamp: time.Now().UnixNano(), Event: e}
	return entry
}

// SetOutput set the output io.Writer for global event logger
func SetOutput(w io.Writer) {
	eventLog.SetOutput(w)
}

// EventLog logs an event defined by caller's struct and output it in JSON format
func LogEvent(event interface{}) {
	var entry *EventLogEntry
	if le, ok := event.(LoggableEvent); ok {
		// allows user to customize the information to be logged
		entry = newEventLogEntry(le.ToLogRecord())
	} else {
		entry = newEventLogEntry(event)
	}
	logJson, e := json.Marshal(entry)
	if e != nil {
		eventLog.Printf("{error: %s}\n", e.Error())
		return
	}

	eventLog.Printf("%s\n", logJson)
}
