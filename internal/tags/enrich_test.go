package tags

import (
	"testing"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func TestEnrich(t *testing.T) {
	txns := []parser.Transaction{{Description: "UBER TRIP", Category: ""}}
	got := Enrich(txns, map[string]string{"UBER": "Transport"})
	if got[0].Category != "Transport" {
		t.Fatalf("cat %q", got[0].Category)
	}
}
