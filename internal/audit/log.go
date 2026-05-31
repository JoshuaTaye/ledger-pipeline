package audit

import (
	"fmt"
	"io"
	"time"
)

// Entry records one pipeline stage event.
type Entry struct {
	Timestamp time.Time
	Stage     string
	Message   string
}

// Log collects audit entries in order.
type Log struct {
	entries []Entry
}

func NewLog() *Log {
	return &Log{}
}

func (l *Log) Record(stage, message string) {
	l.entries = append(l.entries, Entry{
		Timestamp: time.Now().UTC(),
		Stage:     stage,
		Message:   message,
	})
}

func (l *Log) Entries() []Entry {
	out := make([]Entry, len(l.entries))
	copy(out, l.entries)
	return out
}

func (l *Log) Write(w io.Writer) error {
	for _, e := range l.entries {
		if _, err := fmt.Fprintf(w, "%s [%s] %s\n",
			e.Timestamp.Format(time.RFC3339), e.Stage, e.Message); err != nil {
			return err
		}
	}
	return nil
}
