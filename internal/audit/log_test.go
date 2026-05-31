package audit

import (
	"bytes"
	"strings"
	"testing"
)

func TestLog(t *testing.T) {
	log := NewLog()
	log.Record("parse", "loaded 10 rows")
	log.Record("filter", "kept 8 rows")
	if len(log.Entries()) != 2 {
		t.Fatalf("Entries() len = %d", len(log.Entries()))
	}
	var buf bytes.Buffer
	if err := log.Write(&buf); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "[parse]") {
		t.Fatalf("Write() = %q", buf.String())
	}
}
