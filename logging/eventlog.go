package logging

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type EventLogEntry struct {
	id        string
	timestamp int64
	event     interface{}
}

func newEventLogEntry(e interface{}) *EventLogEntry {
	entry := &EventLogEntry{id: uuid.NewV4().String(), timestamp: time.Now().UnixNano(), event: e}
	return entry
}

// EventLog logs an event defined by caller in json format
func EventLog(event interface{}) {
	entry := newEventLogEntry(event)
	logJson, e := json.Marshal(entry)
	if e != nil {
		log.Printf("{error: %s}\n", e.Error())
		return
	}

	log.Printf("%s\n", logJson)

}
